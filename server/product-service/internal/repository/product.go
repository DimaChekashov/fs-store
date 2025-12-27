package repository

import (
	"sync"

	"product-service/internal/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	// GetByID(id string) (*models.Product, error)
	// Create(product models.Product) (*models.Product, error)
	// Update(id string, product models.Product) (*models.Product, error)
	// Delete(id string) error
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

// func GetByID(id string) (*models.Product, error) {

// }

// func Create(product models.Product) (*models.Product, error) {

// }

// func Update(id string, product models.Product) (*models.Product, error) {

// }

// func Delete(id string) error {

// }
