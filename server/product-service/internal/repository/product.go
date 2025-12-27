package repository

import (
	"sync"

	"product-service/internal/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
}

type InMemoryProductRepository struct {
	products map[string]models.Product
	mu       sync.RWMutex
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[string]models.Product),
	}
}

func (r *InMemoryProductRepository) GetAll() ([]models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	products := make([]models.Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}

	return products, nil
}
