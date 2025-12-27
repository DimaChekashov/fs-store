package main

import (
	"log"
	"product-service/internal/handlers"
	"product-service/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewInMemoryProductRepository()

	productHandler := handlers.NewProductHandler(repo)

	r := gin.Default()

	r.GET("/api/products", productHandler.GetAllProducts)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"service": "product-service",
		})
	})

	log.Println("Product service started on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
