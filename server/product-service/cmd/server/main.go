package main

import (
    "github.com/joho/godotenv"
	"log"
	"os"
	"product-service/internal/handlers"
	"product-service/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system environment variables")
    }

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

	port := os.Getenv("SERVER_PORT")
	log.Println("Product service started on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
