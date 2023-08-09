package product

import (
	domain "go-tdd-clean/09/product/domain"
	"time"
)

type ProductRepositoryFake struct {
	storage []domain.Product
	err     error
}

func NewProductRepositoryFake() *ProductRepositoryFake {
	return &ProductRepositoryFake{
		storage: []domain.Product{},
	}
}

func (r *ProductRepositoryFake) Save(p domain.Product) error {
	r.storage = append(r.storage, p)
	return nil
}

func (r *ProductRepositoryFake) Query(fn func(domain.Product) bool) ([]domain.Product, error) {
	result := []domain.Product{}
	for _, item := range r.storage {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result, nil
}

func (r *ProductRepositoryFake) Seed() {
	r.storage = []domain.Product{
		{
			ID:        "1",
			Name:      "Product 1",
			Price:     100,
			CreatedAt: string(time.Now().Format("2006-01-02")),
		},
		{
			ID:        "2",
			Name:      "Product 2",
			Price:     200,
			CreatedAt: string(time.Now().Format("2006-01-02")),
		},
		{
			ID:        "3",
			Name:      "Product 3",
			Price:     300,
			CreatedAt: string(time.Now().Add(-3600 * time.Hour).Format("2006-01-02")),
		},
	}
}
