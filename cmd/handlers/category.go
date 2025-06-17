package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"super-catalog/internal/category"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type SizeRequest struct {
	ID   string `json:"id" validate:"required,max=50"`
	Name string `json:"name" validate:"required,max=100"`
}

type AskGroupOptionRequest struct {
	ID          string `json:"id" validate:"required,max=50"`
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"max=255"`
	Value       int64  `json:"value" validate:"required,min=0"`
}

type AskGroupRequest struct {
	ID           string                  `json:"id" validate:"required,max=50"`
	Group        string                  `json:"group" validate:"required,max=100"`
	MinimunLimit int                     `json:"min_limit" validate:"required,min=0"`
	MaximunLimit int                     `json:"max_limit" validate:"required,min=0"`
	Options      []AskGroupOptionRequest `json:"options" validate:"dive,required"`
}

type SchedulRequest struct {
	Day   string `json:"day" validate:"required,max=20"`
	Hours string `json:"hours" validate:"required,max=20"`
}

type BaseCategoryRequest struct {
	Type        string `json:"type" validate:"required,oneof=FOODS MARKET SCHEDULED"`
	ID          string `json:"id" validate:"required,max=50"`
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"max=255"`
}

type FoodsCategoryRequest struct {
	BaseCategoryRequest
	Culinary  string            `json:"culinary" validate:"required,max=100"`
	Sizes     []SizeRequest     `json:"sizes" validate:"dive,required"`
	AskGroups []AskGroupRequest `json:"ask_groups" validate:"dive"`
}

type MaketCategoryRequest struct {
	BaseCategoryRequest
	Section string `json:"section" validate:"required,max=100"`
}

type ScheduledCategoryRequest struct {
	BaseCategoryRequest
	Schedul []SchedulRequest `json:"schedul" validate:"dive,required"`
}

func validateCategoryRequest(req interface{}) error {
	return validate.Struct(req)
}

func CreateCategoryHandler(c *gin.Context) {
	var rawCategories []map[string]interface{}
	if err := c.ShouldBindJSON(&rawCategories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categories := make([]category.Category, 0, len(rawCategories))
	for _, raw := range rawCategories {
		typeStr, ok := raw["type"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "type field is required"})
			return
		}
		switch category.CategoryType(typeStr) {
		case category.CategoryTypeFoods:
			var req FoodsCategoryRequest
			if err := mapToStruct(raw, &req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := validateCategoryRequest(req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			categories = append(categories, req.ToCategory())
		case category.CategoryTypeMarket:
			var req MaketCategoryRequest
			if err := mapToStruct(raw, &req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := validateCategoryRequest(req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			categories = append(categories, req.ToCategory())
		case category.CategoryTypeScheduled:
			var req ScheduledCategoryRequest
			if err := mapToStruct(raw, &req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := validateCategoryRequest(req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			categories = append(categories, req.ToCategory())
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category type"})
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := category.InsertCategories(ctx, categories); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save categories"})
		return
	}
	c.JSON(http.StatusCreated, categories)
}

func mapToStruct(m map[string]interface{}, out interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

func (cr FoodsCategoryRequest) ToCategory() category.Category {
	sizes := make([]category.Size, len(cr.Sizes))
	for i, s := range cr.Sizes {
		sizes[i] = category.Size{
			ID:   s.ID,
			Name: s.Name,
		}
	}
	askGroups := make([]category.AskGroup, len(cr.AskGroups))
	for i, ag := range cr.AskGroups {
		options := make([]category.AskGroupOption, len(ag.Options))
		for j, o := range ag.Options {
			options[j] = category.AskGroupOption{
				ID:          o.ID,
				Name:        o.Name,
				Description: o.Description,
				Value:       o.Value,
			}
		}
		askGroups[i] = category.AskGroup{
			ID:           ag.ID,
			Group:        ag.Group,
			MinimunLimit: ag.MinimunLimit,
			MaximunLimit: ag.MaximunLimit,
			Options:      options,
		}
	}
	return category.Category{
		Type:        cr.Type,
		ID:          cr.ID,
		Name:        cr.Name,
		Description: cr.Description,
		Culinary:    cr.Culinary,
		Sizes:       sizes,
		AskGroups:   askGroups,
	}
}

func (cr MaketCategoryRequest) ToCategory() category.Category {
	return category.Category{
		Type:        cr.Type,
		ID:          cr.ID,
		Name:        cr.Name,
		Section:     cr.Section,
		Description: cr.Description,
	}
}

func (cr ScheduledCategoryRequest) ToCategory() category.Category {
	schedul := make([]category.Schedul, len(cr.Schedul))
	for i, s := range cr.Schedul {
		schedul[i] = category.Schedul{
			Day:   s.Day,
			Hours: s.Hours,
		}
	}
	return category.Category{
		Type:        cr.Type,
		ID:          cr.ID,
		Name:        cr.Name,
		Description: cr.Description,
		Schedul:     schedul,
	}
}
