package product

import (
	"errors"
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

func (p *Product) Validate() (errs []error) {
	if p.Name == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if p.Price <= 0 {
		errs = append(errs, errors.New("price must be greater than 0"))
	}
	if p.ID == "" {
		errs = append(errs, errors.New("id is required"))
	}
	if p.CreatedAt == "" {
		errs = append(errs, errors.New("created_at is required"))
	}
	return errs
}
