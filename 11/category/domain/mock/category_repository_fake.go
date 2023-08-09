package mock

import (
	"errors"
	category "go-tdd-clean/11/category/domain"
	"time"
)

type CategoryRepositoryFake struct {
	storage []category.Category
}

func NewProductRepositoryFake() *CategoryRepositoryFake {
	return &CategoryRepositoryFake{
		storage: []category.Category{},
	}
}

func (r *CategoryRepositoryFake) Save(p category.Category) error {
	r.storage = append(r.storage, p)
	return nil
}

func (r *CategoryRepositoryFake) GetByName(name string) (category.Category, error) {
	for _, item := range r.storage {
		if item.Name == name {
			return item, nil
		}
	}
	return category.Category{}, errors.New("not found")
}

func (r *CategoryRepositoryFake) Seed() {
	r.storage = []category.Category{
		{
			ID:        "1",
			Name:      "Category A",
			CreatedAt: string(time.Now().Format("2006-01-02")),
		},
		{
			ID:        "2",
			Name:      "Category B",
			CreatedAt: string(time.Now().Format("2006-01-02")),
		},
		{
			ID:        "3",
			Name:      "Category C",
			CreatedAt: string(time.Now().Add(-3600 * time.Hour).Format("2006-01-02")),
		},
	}
}
