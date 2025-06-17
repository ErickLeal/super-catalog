package requests

import (
	"super-catalog/internal/product"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductDetailRequest struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type AdditionalRequest struct {
	ProductID string `json:"product_id"`
}

type UnitRequest struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type BaseProductRequest struct {
	CategoryID        string                 `json:"category_id"`
	ID                string                 `json:"id"`
	Name              string                 `json:"name"`
	Description       string                 `json:"description"`
	Value             int64                  `json:"value"`
	InventoryQuantity int                    `json:"inventory_quantity"`
	IsInventoryActive bool                   `json:"is_inventory_active"`
	ProductDetails    []ProductDetailRequest `json:"product_details"`
}

type ProductFoodsRequest struct {
	BaseProductRequest
	Tags        []string            `json:"tags"`
	Adittionals []AdditionalRequest `json:"adittionals"`
}

type ProductMarketRequest struct {
	BaseProductRequest
	EanCode string      `bson:"ean_code" json:"ean_code"`
	Unit    UnitRequest `json:"unit"`
}

type ProductScheduledRequest struct {
	BaseProductRequest
	FictionalField string `json:"fictional_field"`
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
