package product

import (
	"log"
)

type CreateProductUseCase struct {
	Repository IProductRepository
}

func NewCreateProductUseCase(repository IProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		Repository: repository,
	}
}

func (c *CreateProductUseCase) Execute(input CreateProductInput) (output Output) {
	// validate input
	err := input.Validate()
	if err != nil {
		output.SetError(DomainCodeInvalidInput, err)
		return
	}

	// verify if product already exists
	entity, err := c.Repository.GetByName(input.Name)
	if err != nil {
		log.Println(err)
		output.SetError(DomainCodeInternalError, err)
		return
	}
	if entity != nil {
		output.SetError(DomainCodeAlreadyExists, err)
		return
	}

	// create entity
	entity, err = NewProduct(input.Name, input.Price)
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
	output.SetOk(*entity)
	return
}

type GetProductByMinMaxPrice struct {
	Repository IProductRepository
}

func NewGetProductByMinMaxPriceUseCase(repository IProductRepository) *GetProductByMinMaxPrice {
	return &GetProductByMinMaxPrice{
		Repository: repository,
	}
}

func (c *GetProductByMinMaxPrice) Execute(input GetProductByMinMaxPriceInput) (output Output) {
	// validate input (fail fast)
	err := input.Validate()
	if err != nil {
		output.SetError(DomainCodeInvalidInput, err)
		return
	}

	// query entities
	fn := func(p *Product) bool {
		return p.Price >= input.MinPrice && p.Price <= input.MaxPrice
	}

	entities, err := c.Repository.Query(fn)
	if err != nil {
		log.Println(err)
		output.SetError(DomainCodeInternalError, err)
		return
	}

	// return ok
	output.SetOk(*entities)
	return
}
