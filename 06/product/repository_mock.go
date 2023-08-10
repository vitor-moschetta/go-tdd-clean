package product

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
