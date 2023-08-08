package mock

import (
	"errors"
	product "go-tdd-clean/09/product/domain"
)

type ProductRepositoryFake struct {
	storage []product.Product
	err     error
}

func NewProductRepositoryFake() *ProductRepositoryFake {
	return &ProductRepositoryFake{
		storage: []product.Product{},
	}
}

func (r *ProductRepositoryFake) Save(p product.Product) error {
	if r.err != nil {
		return r.err
	}
	r.storage = append(r.storage, p)
	return nil
}

func (r *ProductRepositoryFake) SetError() {
	r.err = errors.New("error")
}
