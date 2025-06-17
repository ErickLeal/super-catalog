package handlers

import (
	"context"
	"net/http"
	"time"

	"super-catalog/internal/category"

	"github.com/gin-gonic/gin"
)

type CategoryRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (cr CategoryRequest) ToCategory() category.Category {
	return category.Category{
		ID:   cr.ID,
		Name: cr.Name,
	}
}

func CreateCategoryHandler(c *gin.Context) {
	var categoriesReq []CategoryRequest
	if err := c.ShouldBindJSON(&categoriesReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	categories := make([]category.Category, len(categoriesReq))
	for i, req := range categoriesReq {
		categories[i] = req.ToCategory()
	}
	err := category.InsertCategories(ctx, categories)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save categories"})
		return
	}
	c.JSON(http.StatusCreated, categories)
}
