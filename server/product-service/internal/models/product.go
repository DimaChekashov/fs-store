package models

import "time"

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock" binding:"gte=0"`
}

type UpdateProductRequest struct {
	Name        *string  `json:""`
	Description *string  `json:""`
	Price       *float64 `json:"" binding:"omitempty,gt=0"`
	Category    *string  `json:""`
	Stock       *int     `json:"" binding:"omitempty,gte=0"`
}
