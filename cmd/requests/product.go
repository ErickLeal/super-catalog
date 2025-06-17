package requests

import (
	"super-catalog/internal/product"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductDetailRequest struct {
	Name  string `json:"name" validate:"required,max=100"`
	Value int64  `json:"value" validate:"required,min=0"`
}

type AdditionalRequest struct {
	ProductID string `json:"product_id" validate:"required,max=50"`
}

type UnitRequest struct {
	Name  string `json:"name" validate:"required,max=100"`
	Value int64  `json:"value" validate:"required,min=0"`
}

type BaseProductRequest struct {
	CategoryID        string                 `json:"category_id" validate:"required,max=50"`
	ID                string                 `json:"id" validate:"required,max=50"`
	Name              string                 `json:"name" validate:"required,max=100"`
	Description       string                 `json:"description" validate:"max=255"`
	Value             int64                  `json:"value" validate:"required,min=0"`
	InventoryQuantity int                    `json:"inventory_quantity" validate:"required,min=0"`
	IsInventoryActive bool                   `json:"is_inventory_active" validate:"required"`
	ProductDetails    []ProductDetailRequest `json:"product_details" validate:"dive"`
}

type ProductFoodsRequest struct {
	BaseProductRequest
	Tags        []string            `json:"tags" validate:"required,dive,max=50"`
	Adittionals []AdditionalRequest `json:"adittionals" validate:"dive"`
}

type ProductMarketRequest struct {
	BaseProductRequest
	EanCode string      `bson:"ean_code" json:"ean_code" validate:"required,max=50"`
	Unit    UnitRequest `json:"unit" validate:"required"`
}

type ProductScheduledRequest struct {
	BaseProductRequest
	FictionalField string `json:"fictional_field" validate:"required,max=100"`
}

func (r ProductDetailRequest) ToModel() product.ProductDetail {
	return product.ProductDetail{
		Name:  r.Name,
		Value: r.Value,
	}
}

func (r AdditionalRequest) ToModel() product.Adittional {
	return product.Adittional{
		ProductID: r.ProductID,
	}
}

func (r UnitRequest) ToModel() product.Unit {
	return product.Unit{
		Name:  r.Name,
		Value: r.Value,
	}
}

func (r ProductFoodsRequest) ToModel(cat map[string]interface{}) product.ProductFoods {

	productDetails := make([]product.ProductDetail, len(r.ProductDetails))
	for i, d := range r.ProductDetails {
		productDetails[i] = d.ToModel()
	}
	adittionals := make([]product.Adittional, len(r.Adittionals))
	for i, a := range r.Adittionals {
		adittionals[i] = a.ToModel()
	}
	return product.ProductFoods{
		Category: product.CategoryProduct{
			MongoId: cat["_id"].(primitive.ObjectID),
			ID:      cat["id"].(string),
			Type:    cat["type"].(string),
		},
		ID:                r.ID,
		Name:              r.Name,
		Description:       r.Description,
		Value:             r.Value,
		InventoryQuantity: r.InventoryQuantity,
		IsInventoryActive: r.IsInventoryActive,
		ProductDetails:    productDetails,
		Tags:              r.Tags,
		Adittionals:       adittionals,
	}
}

func (r ProductMarketRequest) ToModel(cat map[string]interface{}) product.ProductMarket {
	productDetails := make([]product.ProductDetail, len(r.ProductDetails))
	for i, d := range r.ProductDetails {
		productDetails[i] = d.ToModel()
	}

	return product.ProductMarket{
		Category: product.CategoryProduct{
			MongoId: cat["_id"].(primitive.ObjectID),
			ID:      cat["id"].(string),
			Type:    cat["type"].(string),
		},
		ID:                r.ID,
		Name:              r.Name,
		Description:       r.Description,
		Value:             r.Value,
		InventoryQuantity: r.InventoryQuantity,
		IsInventoryActive: r.IsInventoryActive,
		ProductDetails:    productDetails,
		EanCode:           r.EanCode,
		Unit:              r.Unit.ToModel(),
	}
}

func (r ProductScheduledRequest) ToModel(cat map[string]interface{}) product.ProductScheduled {
	productDetails := make([]product.ProductDetail, len(r.ProductDetails))
	for i, d := range r.ProductDetails {
		productDetails[i] = d.ToModel()
	}

	return product.ProductScheduled{
		Category: product.CategoryProduct{
			MongoId: cat["_id"].(primitive.ObjectID),
			ID:      cat["id"].(string),
			Type:    cat["type"].(string),
		},
		ID:                r.ID,
		Name:              r.Name,
		Description:       r.Description,
		Value:             r.Value,
		InventoryQuantity: r.InventoryQuantity,
		IsInventoryActive: r.IsInventoryActive,
		ProductDetails:    productDetails,
		FictionalField:    r.FictionalField,
	}
}
