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
