package handlers

import (
	"context"
	"net/http"
	"time"

	"super-catalog/cmd/helpers"
	"super-catalog/cmd/requests"
	"super-catalog/internal/category"

	"github.com/gin-gonic/gin"
)

type categoryRequestHandler struct {
	Type      category.CategoryType
	Unmarshal func(map[string]interface{}) (interface{}, error)
	Validate  func(interface{}) error
	ToModel   func(interface{}) interface{}
}

var categoryRequestHandlers = []categoryRequestHandler{
	{
		Type: category.CategoryTypeFoods,
		Unmarshal: func(raw map[string]interface{}) (interface{}, error) {
			var req requests.FoodsCategoryRequest
			if err := helpers.MapToStruct(raw, &req); err != nil {
				return nil, err
			}
			return req, nil
		},
		Validate: func(req interface{}) error {
			return helpers.ValidateRequest(req)
		},
		ToModel: func(req interface{}) interface{} {
			return req.(requests.FoodsCategoryRequest).ToCategory()
		},
	},
	{
		Type: category.CategorySlicedFoods,
		Unmarshal: func(raw map[string]interface{}) (interface{}, error) {
			var req requests.SlicedFoodsCategoryRequest
			if err := helpers.MapToStruct(raw, &req); err != nil {
				return nil, err
			}
			return req, nil
		},
		Validate: func(req interface{}) error {
			return helpers.ValidateRequest(req)
		},
		ToModel: func(req interface{}) interface{} {
			return req.(requests.SlicedFoodsCategoryRequest).ToCategory()
		},
	},
	{
		Type: category.CategoryTypeMarket,
		Unmarshal: func(raw map[string]interface{}) (interface{}, error) {
			var req requests.MaketCategoryRequest
			if err := helpers.MapToStruct(raw, &req); err != nil {
				return nil, err
			}
			return req, nil
		},
		Validate: func(req interface{}) error {
			return helpers.ValidateRequest(req)
		},
		ToModel: func(req interface{}) interface{} {
			return req.(requests.MaketCategoryRequest).ToCategory()
		},
	},
	{
		Type: category.CategoryTypeScheduled,
		Unmarshal: func(raw map[string]interface{}) (interface{}, error) {
			var req requests.ScheduledCategoryRequest
			if err := helpers.MapToStruct(raw, &req); err != nil {
				return nil, err
			}
			return req, nil
		},
		Validate: func(req interface{}) error {
			return helpers.ValidateRequest(req)
		},
		ToModel: func(req interface{}) interface{} {
			return req.(requests.ScheduledCategoryRequest).ToCategory()
		},
	},
	{
		Type: category.CategoryTypeOpen,
		Unmarshal: func(raw map[string]interface{}) (interface{}, error) {
			var req requests.OpenCategoryRequest
			if err := helpers.MapToStruct(raw, &req); err != nil {
				return nil, err
			}
			return req, nil
		},
		Validate: func(req interface{}) error {
			return helpers.ValidateRequest(req)
		},
		ToModel: func(req interface{}) interface{} {
			return req.(requests.OpenCategoryRequest).ToCategory()
		},
	},
}

func CreateCategoryHandler(c *gin.Context) {
	var rawCategories []map[string]interface{}
	if err := c.ShouldBindJSON(&rawCategories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categories := make([]interface{}, 0, len(rawCategories))
	for i, raw := range rawCategories {
		typeStr, ok := raw["type"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "type field is required", "index": i})
			return
		}
		handler := getCategoryRequestHandler(typeStr)
		if handler == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category type", "type": typeStr, "index": i})
			return
		}
		req, err := handler.Unmarshal(raw)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "index": i})
			return
		}
		if err := handler.Validate(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "index": i})
			return
		}
		categories = append(categories, handler.ToModel(req))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := category.InsertCategories(ctx, categories); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save categories", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, categories)
}

func getCategoryRequestHandler(typeStr string) *categoryRequestHandler {
	for _, h := range categoryRequestHandlers {
		if string(h.Type) == typeStr {
			return &h
		}
	}
	return nil
}
