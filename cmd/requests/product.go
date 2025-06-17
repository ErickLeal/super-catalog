package requests

type ProductDetailRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AdditionalRequest struct {
	ProductID string `json:"product_id"`
}

type UnitRequest struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type BaseProductRequest struct {
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
