package category

type CategoryRepositoryInMemory struct {
	storage []Category
	err     error
}

func NewCategoryRepositoryInMemory() *CategoryRepositoryInMemory {
	return &CategoryRepositoryInMemory{
		storage: []Category{},
	}
}

func (r *CategoryRepositoryInMemory) Save(p Category) error {
	r.storage = append(r.storage, p)
	return nil
}

func (r *CategoryRepositoryInMemory) GetByID(id string) (Category, error) {
	for _, item := range r.storage {
		if item.ID == id {
			return item, nil
		}
	}
	return Category{}, nil
}
