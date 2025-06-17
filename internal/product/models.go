package product

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryProduct struct {
	MongoId primitive.ObjectID `bson:"_id" json:"-"`
	ID      string             `bson:"id" json:"id"`
	Type    string             `bson:"type" json:"type"`
}

type Adittional struct {
	ProductID string `bson:"product_id" json:"product_id"`
}

type Unit struct {
	Name  string `bson:"name" json:"name"`
	Value int64  `bson:"value" json:"value"`
}

type ProductDetail struct {
	Name  string `bson:"name" json:"name"`
	Value int64  `bson:"value" json:"value"`
}

type ProductFoods struct {
	Category          CategoryProduct `bson:"category" json:"category"`
	ID                string          `bson:"id" json:"id"`
	Name              string          `bson:"name" json:"name"`
	Description       string          `bson:"description" json:"description"`
	Value             int64           `bson:"value" json:"value"`
	InventoryQuantity int             `bson:"inventory_quantity" json:"inventory_quantity"`
	IsInventoryActive bool            `bson:"is_inventory_active" json:"is_inventory_active"`
	ProductDetails    []ProductDetail `bson:"product_details" json:"product_details"`
	Tags              []string        `bson:"tags" json:"tags"`
	Adittionals       []Adittional    `bson:"adittionals" json:"adittionals"`
}

type ProductMarket struct {
	Category          CategoryProduct `bson:"category" json:"category"`
	ID                string          `bson:"id" json:"id"`
	Name              string          `bson:"name" json:"name"`
	Description       string          `bson:"description" json:"description"`
	Value             int64           `bson:"value" json:"value"`
	InventoryQuantity int             `bson:"inventory_quantity" json:"inventory_quantity"`
	IsInventoryActive bool            `bson:"is_inventory_active" json:"is_inventory_active"`
	ProductDetails    []ProductDetail `bson:"product_details" json:"product_details"`
	EanCode           string          `bson:"ean_code" json:"ean_code"`
	Unit              Unit            `bson:"unit" json:"unit"`
}

type ProductScheduled struct {
	Category          CategoryProduct `bson:"category" json:"category"`
	ID                string          `bson:"id" json:"id"`
	Name              string          `bson:"name" json:"name"`
	Description       string          `bson:"description" json:"description"`
	Value             int64           `bson:"value" json:"value"`
	InventoryQuantity int             `bson:"inventory_quantity" json:"inventory_quantity"`
	IsInventoryActive bool            `bson:"is_inventory_active" json:"is_inventory_active"`
	ProductDetails    []ProductDetail `bson:"product_details" json:"product_details"`
	FictionalField    string          `bson:"fictional_field" json:"fictional_field"` // Placeholder for future use
}
