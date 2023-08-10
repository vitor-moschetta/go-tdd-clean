package product

import (
	"time"
)

type ProductRepositoryInMemory struct {
	storage []Product
}

func NewProductRepositoryInMemory() *ProductRepositoryInMemory {
	return &ProductRepositoryInMemory{
		storage: []Product{},
	}
}

func (r *ProductRepositoryInMemory) Save(p Product) error {
	r.storage = append(r.storage, p)
	return nil
}

func (r *ProductRepositoryInMemory) GetByName(name string) (Product, error) {
	for _, item := range r.storage {
		if item.Name == name {
			return item, nil
		}
	}
	return Product{}, nil
}

func (r *ProductRepositoryInMemory) Query(fn func(Product) bool) ([]Product, error) {
	result := []Product{}
	for _, item := range r.storage {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result, nil
}

func (r *ProductRepositoryInMemory) Seed() {
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
