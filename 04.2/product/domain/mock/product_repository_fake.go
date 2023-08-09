package mock

import (
	product "go-tdd-clean/04.2/product/domain"
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
