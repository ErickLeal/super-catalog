package handlers

import (
	"context"
	"net/http"
	"time"

	"super-catalog/internal/product"

	"github.com/gin-gonic/gin"
)

type ProductDetailRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProductRequest struct {
	ID                string                 `json:"id"`
	Name              string                 `json:"name"`
	Description       string                 `json:"description"`
	Enabled           bool                   `json:"enabled"`
	SKU               string                 `json:"sku"`
	Value             int64                  `json:"value"`
	PromotionalValue  int64                  `json:"promotional_value"`
	InventoryQuantity int                    `json:"inventory_quantity"`
	IsInventoryActive bool                   `json:"is_inventory_active"`
	ImagesURL         []string               `json:"images_url"`
	ProductDetails    []ProductDetailRequest `json:"product_details"`
}

func (pr ProductRequest) ToProduct() product.Product {
	productDetails := make([]product.ProductDetail, len(pr.ProductDetails))
	for i, d := range pr.ProductDetails {
		productDetails[i] = product.ProductDetail{
			Name:  d.Name,
			Value: d.Value,
		}
	}
	return product.Product{
		ID:                pr.ID,
		Name:              pr.Name,
		Description:       pr.Description,
		Enabled:           pr.Enabled,
		SKU:               pr.SKU,
		Value:             pr.Value,
		PromotionalValue:  pr.PromotionalValue,
		InventoryQuantity: pr.InventoryQuantity,
		IsInventoryActive: pr.IsInventoryActive,
		ImagesURL:         pr.ImagesURL,
		ProductDetails:    productDetails,
	}
}

func CreateProductHandler(c *gin.Context) {
	var productsReq []ProductRequest
	if err := c.ShouldBindJSON(&productsReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	products := make([]product.Product, len(productsReq))
	for i, req := range productsReq {
		products[i] = req.ToProduct()
	}
	err := product.InsertProducts(ctx, products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save products"})
		return
	}
	c.JSON(http.StatusCreated, products)
}
