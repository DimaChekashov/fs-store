package handlers

import (
	"net/http"

	"product-service/internal/repository"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  products,
		"count": len(products),
	})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {

}

func (h *ProductHandler) CreateProduct(c *gin.Context) {

}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {

}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {

}
