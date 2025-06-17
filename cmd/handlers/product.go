package handlers

import (
	"context"
	"net/http"
	"super-catalog/cmd/helpers"
	"super-catalog/cmd/requests"
	"super-catalog/internal/category"
	"super-catalog/internal/product"
	"time"

	"github.com/gin-gonic/gin"
)

type productRequestHandler struct {
	Type      string
	Unmarshal func(map[string]interface{}) (interface{}, error)
	Validate  func(interface{}) error
	ToModel   func(interface{}, map[string]interface{}) interface{}
}

var productRequestHandlers = []productRequestHandler{
	{
		Type: string(category.CategoryTypeFoods),
		Unmarshal: func(raw map[string]interface{}) (interface{}, error) {
			var req requests.ProductFoodsRequest
			if err := helpers.MapToStruct(raw, &req); err != nil {
				return nil, err
			}
			return req, nil
		},
		Validate: func(req interface{}) error {
			return helpers.ValidateRequest(req)
		},
		ToModel: func(req interface{}, cat map[string]interface{}) interface{} {
			return req.(requests.ProductFoodsRequest).ToModel(cat)
		},
	},
	{
		Type: string(category.CategoryTypeMarket),
		Unmarshal: func(raw map[string]interface{}) (interface{}, error) {
			var req requests.ProductMarketRequest
			if err := helpers.MapToStruct(raw, &req); err != nil {
				return nil, err
			}
			return req, nil
		},
		Validate: func(req interface{}) error {
			return helpers.ValidateRequest(req)
		},
		ToModel: func(req interface{}, cat map[string]interface{}) interface{} {
			return req.(requests.ProductMarketRequest).ToModel(cat)
		},
	},
	{
		Type: string(category.CategoryTypeScheduled),
		Unmarshal: func(raw map[string]interface{}) (interface{}, error) {
			var req requests.ProductScheduledRequest
			if err := helpers.MapToStruct(raw, &req); err != nil {
				return nil, err
			}
			return req, nil
		},
		Validate: func(req interface{}) error {
			return helpers.ValidateRequest(req)
		},
		ToModel: func(req interface{}, cat map[string]interface{}) interface{} {
			return req.(requests.ProductScheduledRequest).ToModel(cat)
		},
	},
}

func CreateProductHandler(c *gin.Context) {
	var rawProducts []map[string]interface{}
	if err := c.ShouldBindJSON(&rawProducts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products := make([]interface{}, 0, len(rawProducts))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for i, raw := range rawProducts {
		categoryID, ok := raw["category_id"].(string)
		if !ok || categoryID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "category_id field is required", "index": i})
			return
		}
		cat, err := category.GetCategoryByID(ctx, categoryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "category not found", "category_id": categoryID, "index": i})
			return
		}
		typeStr, ok := cat["type"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "category type not found", "category_id": categoryID, "index": i})
			return
		}
		handler := getProductRequestHandler(typeStr)
		if handler == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product type for category", "type": typeStr, "index": i})
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

		products = append(products, handler.ToModel(req, cat))
	}

	if err := product.InsertProducts(ctx, products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save products", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, products)
}

func getProductRequestHandler(typeStr string) *productRequestHandler {
	for _, h := range productRequestHandlers {
		if h.Type == typeStr {
			return &h
		}
	}
	return nil
}
