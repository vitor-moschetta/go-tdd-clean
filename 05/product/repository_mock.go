package product

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
