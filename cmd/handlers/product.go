package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func processProduct(ctx context.Context, raw map[string]interface{}, i int, c *gin.Context) (interface{}, bool) {
	time.Sleep(100 * time.Millisecond) // Simula processamento
	categoryID, ok := raw["category_id"].(string)
	if !ok || categoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category_id field is required", "index": i})
		return nil, false
	}
	cat, err := category.GetCategoryByID(ctx, categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category not found", "category_id": categoryID, "index": i})
		return nil, false
	}
	typeStr, ok := cat["type"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category type not found", "category_id": categoryID, "index": i})
		return nil, false
	}
	handler := getProductRequestHandler(typeStr)
	if handler == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product type for category", "type": typeStr, "index": i})
		return nil, false
	}
	req, err := handler.Unmarshal(raw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "index": i})
		return nil, false
	}
	if err := handler.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "index": i})
		return nil, false
	}
	return handler.ToModel(req, cat), true
}

func CreateProductHandler(c *gin.Context) {
	var rawProducts []map[string]interface{}
	if err := c.ShouldBindJSON(&rawProducts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products := make([]interface{}, 0, len(rawProducts))
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	for i, raw := range rawProducts {
		product, ok := processProduct(ctx, raw, i, c)
		if !ok {
			return
		}
		products = append(products, product)
	}

	if err := product.InsertProducts(ctx, products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save products", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, products)
}

// Gera 100 registros de cada tipo e salva em examples/products_100_each_type.json
func GenerateProductsFileHandler(c *gin.Context) {
	type ProductScheduled struct {
		CategoryID        string                   `json:"category_id"`
		ID                string                   `json:"id"`
		Name              string                   `json:"name"`
		Description       string                   `json:"description"`
		Value             int                      `json:"value"`
		InventoryQuantity int                      `json:"inventory_quantity"`
		IsInventoryActive bool                     `json:"is_inventory_active"`
		ProductDetails    []map[string]interface{} `json:"product_details"`
		FictionalField    string                   `json:"fictional_field"`
	}

	type ProductMarket struct {
		CategoryID        string                   `json:"category_id"`
		ID                string                   `json:"id"`
		Name              string                   `json:"name"`
		Description       string                   `json:"description"`
		Value             int                      `json:"value"`
		InventoryQuantity int                      `json:"inventory_quantity"`
		IsInventoryActive bool                     `json:"is_inventory_active"`
		ProductDetails    []map[string]interface{} `json:"product_details"`
		EanCode           string                   `json:"ean_code"`
		Unit              map[string]interface{}   `json:"unit"`
	}

	type ProductFoods struct {
		CategoryID        string                   `json:"category_id"`
		ID                string                   `json:"id"`
		Name              string                   `json:"name"`
		Description       string                   `json:"description"`
		Value             int                      `json:"value"`
		InventoryQuantity int                      `json:"inventory_quantity"`
		IsInventoryActive bool                     `json:"is_inventory_active"`
		Tags              []string                 `json:"tags"`
		Adittionals       []map[string]interface{} `json:"adittionals"`
	}

	var products []interface{}
	// SCHEDULED
	for i := 1; i <= 100; i++ {
		products = append(products, ProductScheduled{
			CategoryID:        "2",
			ID:                fmt.Sprintf("prod-%03d", i),
			Name:              fmt.Sprintf("Agendamento de Sala %d", i),
			Description:       fmt.Sprintf("Reserva de sala para reuniões %d", i),
			Value:             15000 + i*5000,
			InventoryQuantity: i,
			IsInventoryActive: true,
			ProductDetails: []map[string]interface{}{
				{"name": "sala", "value": 100 + i},
				{"name": "andar", "value": i/10 + 1},
			},
			FictionalField: fmt.Sprintf("Agendamento especial %d", i),
		})
	}
	// MARKET
	for i := 1; i <= 100; i++ {
		products = append(products, ProductMarket{
			CategoryID:        "1",
			ID:                fmt.Sprintf("PROD123-%d", i),
			Name:              fmt.Sprintf("Arroz 5kg %d", i),
			Description:       fmt.Sprintf("Arroz branco tipo 1 pacote 5kg %d", i),
			Value:             2599 + i*100,
			InventoryQuantity: 100 + i,
			IsInventoryActive: true,
			ProductDetails: []map[string]interface{}{
				{"name": "Marca", "value": i},
				{"name": "Origem", "value": i + 1},
			},
			EanCode: fmt.Sprintf("7891234567%04d", 890+i),
			Unit:    map[string]interface{}{"name": "UN", "value": 1},
		})
	}
	// FOODS
	for i := 1; i <= 100; i++ {
		products = append(products, ProductFoods{
			CategoryID:        "3",
			ID:                fmt.Sprintf("PROD456-%d", i),
			Name:              fmt.Sprintf("Cachorro quente %d", i),
			Description:       fmt.Sprintf("Cachorro quente tradicional %d", i),
			Value:             4990 + i*100,
			InventoryQuantity: 20 + i,
			IsInventoryActive: true,
			Tags:              []string{"SALGADA"},
			Adittionals: []map[string]interface{}{
				{"product_id": "ADICIONAL1"},
				{"product_id": "ADICIONAL2"},
			},
		})
	}
	filePath := "examples/products_100_each_type.json"
	f, err := os.Create(filePath)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(products); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Arquivo gerado com sucesso", "file": filePath})
}

func getProductRequestHandler(typeStr string) *productRequestHandler {
	for _, h := range productRequestHandlers {
		if h.Type == typeStr {
			return &h
		}
	}
	return nil
}
