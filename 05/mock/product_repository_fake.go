package mock

import (
	"go-tdd-clean/05/product"
)

type ProductRepositoryFake struct {
	Storage []product.Product
}

func NewProductRepositoryFake() *ProductRepositoryFake {
	return &ProductRepositoryFake{
		Storage: []product.Product{},
	}
}

func (r *ProductRepositoryFake) Save(p product.Product) error {
	r.Storage = append(r.Storage, p)
	return nil
}
