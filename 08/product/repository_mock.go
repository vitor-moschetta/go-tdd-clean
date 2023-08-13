package product

import (
	"time"
)

type InMemoryProductRepository struct {
	storage []Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		storage: []Product{},
	}
}

func (r *InMemoryProductRepository) Save(p *Product) error {
	r.storage = append(r.storage, *p)
	return nil
}

func (r *InMemoryProductRepository) GetByName(name string) (*Product, error) {
	for _, item := range r.storage {
		if item.Name == name {
			return &item, nil
		}
	}
	return nil, nil
}

func (r *InMemoryProductRepository) Query(fn func(*Product) bool) (*[]Product, error) {
	result := []Product{}
	for _, item := range r.storage {
		if fn(&item) {
			result = append(result, item)
		}
	}
	return &result, nil
}

func (r *InMemoryProductRepository) Seed() {
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
