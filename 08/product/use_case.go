package product

import (
	"log"
)

type ProductUseCase struct {
	Repository IProductRepository
}

func NewProductUseCase(repository IProductRepository) *ProductUseCase {
	return &ProductUseCase{
		Repository: repository,
	}
}

func (c *ProductUseCase) Create(input CreateProductInput) (output Output) {
	// validate input
	err := input.Validate()
	if err != nil {
		output.SetError(DomainCodeInvalidInput, err)
		return
	}

	// create entity
	entity, err := NewProduct(input.Name, input.Price)
	if err != nil {
		log.Println(err)
		output.SetError(DomainCodeInternalError, err)
		return
	}

	// save entity to storage
	err = c.Repository.Save(entity)
	if err != nil {
		log.Println(err)
		output.SetError(DomainCodeInternalError, err)
		return
	}

	// return ok
	output.SetOk(entity)
	return
}

func (c *ProductUseCase) GetByMinMaxPrice(input GetProductByMinMaxPriceInput) (output Output) {
	// validate input (fail fast)
	err := input.Validate()
	if err != nil {
		output.SetError(DomainCodeInvalidInput, err)
		return
	}

	// query entities
	fn := func(p Product) bool {
		return p.Price >= input.MinPrice && p.Price <= input.MaxPrice
	}

	entities, err := c.Repository.Query(fn)
	if err != nil {
		log.Println(err)
		output.SetError(DomainCodeInternalError, err)
		return
	}

	// return ok
	output.SetOk(entities)
	return
}
