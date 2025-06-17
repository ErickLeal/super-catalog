package handlers

import (
	"context"
	"net/http"
	"time"

	"super-catalog/internal/product"

	"github.com/gin-gonic/gin"
)

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
