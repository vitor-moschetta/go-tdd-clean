package product

import (
	"time"
)

type ProductRepositoryFake struct {
	storage []Product
}

func NewProductRepositoryFake() *ProductRepositoryFake {
	return &ProductRepositoryFake{
		storage: []Product{},
	}
}

func (r *ProductRepositoryFake) Save(p Product) error {
	r.storage = append(r.storage, p)
	return nil
}

func (r *ProductRepositoryFake) Query(fn func(Product) bool) ([]Product, error) {
	result := []Product{}
	for _, item := range r.storage {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result, nil
}

func (r *ProductRepositoryFake) Seed() {
	r.storage = []Product{
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
