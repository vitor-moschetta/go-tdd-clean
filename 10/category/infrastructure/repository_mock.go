package category

import category "go-tdd-clean/10/category/domain"

type CategoryRepositoryFake struct {
	storage []category.Category
	err     error
}

func NewCategoryRepositoryFake() *CategoryRepositoryFake {
	return &CategoryRepositoryFake{
		storage: []category.Category{},
	}
}

func (r *CategoryRepositoryFake) Save(p category.Category) error {
	r.storage = append(r.storage, p)
	return nil
}

func (r *CategoryRepositoryFake) GetByID(id string) (category.Category, error) {
	for _, item := range r.storage {
		if item.ID == id {
			return item, nil
		}
	}
	return category.Category{}, nil
}
