package product

import (
	"context"
	"super-catalog/internal/infra"
)

type ProductDetail struct {
	Name  string `bson:"name" json:"name"`
	Value string `bson:"value" json:"value"`
}

type Product struct {
	ID                string          `bson:"id,omitempty" json:"id"`
	Name              string          `bson:"name" json:"name"`
	Description       string          `bson:"description" json:"description"`
	Enabled           bool            `bson:"enabled" json:"enabled"`
	SKU               string          `bson:"sku" json:"sku"`
	Value             int64           `bson:"value" json:"value"`
	PromotionalValue  int64           `bson:"promotional_value" json:"promotional_value"`
	InventoryQuantity int             `bson:"inventory_quantity" json:"inventory_quantity"`
	IsInventoryActive bool            `bson:"is_inventory_active" json:"is_inventory_active"`
	ImagesURL         []string        `bson:"images_url" json:"images_url"`
	ProductDetails    []ProductDetail `bson:"product_details" json:"product_details"`
}

func InsertProducts(ctx context.Context, products []Product) error {
	coll, err := infra.GetCollection("supercatalog", "products")
	if err != nil {
		return err
	}
	insertDocs := make([]interface{}, len(products))
	for i, p := range products {
		insertDocs[i] = p
	}
	_, err = coll.InsertMany(ctx, insertDocs)
	return err
}
