package mock

import (
	"go-tdd-clean/04/product"
)

type ProductRepositoryFake struct {
	storage []product.Product
}

func NewProductRepositoryFake() *ProductRepositoryFake {
	return &ProductRepositoryFake{
		storage: []product.Product{},
	}
}

func (r *ProductRepositoryFake) Save(p product.Product) error {
	r.storage = append(r.storage, p)
	return nil
}
