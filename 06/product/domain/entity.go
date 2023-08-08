package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        string
	Name      string
	Price     float64
	CreatedAt string
}

func NewProduct(name string, price float64) Product {
	return Product{
		ID:        uuid.New().String(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().Format(time.RFC3339), // 2021-01-01 00:00:00
	}
}

func (p *Product) Validate() bool {
	if p.Name == "" {
		return false
	}
	if p.Price <= 0 {
		return false
	}
	if p.ID == "" {
		return false
	}
	if p.CreatedAt == "" {
		return false
	}
	return true
}
