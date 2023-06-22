package mock

import (
	"errors"
	"go-tdd-clean/10/domain/product"
)

type ProductRepositoryFake struct {
	Storage []product.Product
	err     error
}

func NewProductRepositoryFake() *ProductRepositoryFake {
	return &ProductRepositoryFake{
		Storage: []product.Product{},
	}
}

func (r *ProductRepositoryFake) Save(p product.Product) error {
	if r.err != nil {
		return r.err
	}
	r.Storage = append(r.Storage, p)
	return nil
}

func (r *ProductRepositoryFake) SetError() {
	r.err = errors.New("error")
}
