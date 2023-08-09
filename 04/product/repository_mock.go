package product

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
