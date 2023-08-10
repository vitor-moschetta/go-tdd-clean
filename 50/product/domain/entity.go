package product

import (
	"go-tdd-clean/50/shared"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         string
	Name       string
	Price      float64
	CreatedAt  string
	CategoryID string
}

func NewProduct(name string, price float64, categoryID string) (Product, error) {
	product := Product{
		ID:         uuid.New().String(),
		Name:       name,
		Price:      price,
		CategoryID: categoryID,
		CreatedAt:  time.Now().Format(time.RFC3339), // 2021-01-01 00:00:00
	}
	err := product.validate()
	return product, err
}

func (p *Product) validate() error {
	err := new(shared.Error)
	if p.Name == "" {
		err.AddError("name is required")
	}
	if p.Price <= 0 {
		err.AddError("price is required")
	}
	if p.ID == "" {
		err.AddError("id is required")
	}
	if p.CreatedAt == "" {
		err.AddError("created_at is required")
	}
	if p.CategoryID == "" {
		err.AddError("category_id is required")
	}
	if err.Error() != "" {
		return err
	}
	return nil
}
