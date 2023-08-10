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

func (c *ProductUseCase) Execute(input CreateProductInput) (output Output) {
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
	if entity.ID != "" {
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
	output.SetOk(entity)
	return
}
