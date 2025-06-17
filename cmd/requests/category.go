package requests

import (
	"super-catalog/internal/category"
)

type SizeRequest struct {
	ID   string `json:"id" validate:"required,max=50"`
	Name string `json:"name" validate:"required,max=100"`
}

type SizeFlavorRequest struct {
	ID          string `json:"id" validate:"required,max=50"`
	Name        string `json:"name" validate:"required,max=100"`
	MaxFlavours int64  `json:"max_flavors" validate:"required,max=64,min=1"`
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
	Type        string `json:"type" validate:"required,oneof=FOODS MARKET SCHEDULED SLICED_FOODS"`
	ID          string `json:"id" validate:"required,max=50"`
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"max=255"`
}

type FoodsCategoryRequest struct {
	BaseCategoryRequest
	Culinary  string            `json:"culinary" validate:"required,max=100"`
	StoreId   string            `json:"store_id" validate:"required"`
	Sizes     []SizeRequest     `json:"sizes" validate:"dive,required"`
	AskGroups []AskGroupRequest `json:"ask_groups" validate:"dive"`
}

type SlicedFoodsCategoryRequest struct {
	BaseCategoryRequest
	StoreId   string              `json:"store_id" validate:"required"`
	Sizes     []SizeFlavorRequest `json:"sizes" validate:"dive,required"`
	AskGroups []AskGroupRequest   `json:"ask_groups" validate:"dive"`
}

type MaketCategoryRequest struct {
	BaseCategoryRequest
	Section string `json:"section" validate:"required,max=100"`
}

type ScheduledCategoryRequest struct {
	BaseCategoryRequest
	Schedul []SchedulRequest `json:"schedul" validate:"dive,required"`
}

func (cr FoodsCategoryRequest) ToCategory() category.FoodsCategory {
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
	return category.FoodsCategory{
		Type:        cr.Type,
		StoreId:     cr.StoreId,
		ID:          cr.ID,
		Name:        cr.Name,
		Description: cr.Description,
		Culinary:    cr.Culinary,
		Sizes:       sizes,
		AskGroups:   askGroups,
	}
}

func (cr MaketCategoryRequest) ToCategory() category.MaketCategory {
	return category.MaketCategory{
		Type:        cr.Type,
		ID:          cr.ID,
		Name:        cr.Name,
		Section:     cr.Section,
		Description: cr.Description,
	}
}

func (cr ScheduledCategoryRequest) ToCategory() category.SchedulCategory {
	schedul := make([]category.Schedul, len(cr.Schedul))
	for i, s := range cr.Schedul {
		schedul[i] = category.Schedul{
			Day:   s.Day,
			Hours: s.Hours,
		}
	}
	return category.SchedulCategory{
		Type:        cr.Type,
		ID:          cr.ID,
		Name:        cr.Name,
		Description: cr.Description,
		Schedul:     schedul,
	}
}

func (cr SlicedFoodsCategoryRequest) ToCategory() category.SlicedFoodsCategory {
	sizes := make([]category.SizeFlavor, len(cr.Sizes))
	for i, s := range cr.Sizes {
		sizes[i] = category.SizeFlavor{
			ID:          s.ID,
			Name:        s.Name,
			MaxFlavours: s.MaxFlavours,
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
	return category.SlicedFoodsCategory{
		Type:        cr.Type,
		StoreId:     cr.StoreId,
		ID:          cr.ID,
		Name:        cr.Name,
		Description: cr.Description,
		Sizes:       sizes,
		AskGroups:   askGroups,
	}
}
