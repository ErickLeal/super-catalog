package main

import (
	"github.com/gin-gonic/gin"

	handlers "super-catalog/cmd/handlers"
)

func main() {
	r := gin.Default()
	r.POST("/products", handlers.CreateProductHandler)
	r.POST("/categories", handlers.CreateCategoryHandler)
	r.Run(":8080")
}
