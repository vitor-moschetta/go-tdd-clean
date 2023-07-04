package product

import (
	"go-tdd-clean/10/domain/product"
	"go-tdd-clean/10/shared"
	"log"
)

type ProductUseCase struct {
	Repository product.IProductRepository
}

func NewProductUseCase(repository product.IProductRepository) *ProductUseCase {
	return &ProductUseCase{
		Repository: repository,
	}
}

func (c *ProductUseCase) Create(input CreateProductInput) (output shared.Output) {
	// Validate input (fail fast)
	errs := input.Validate()
	if errs != nil {
		output.SetErrors(shared.DomainCodeInvalidInput, errs)
		return
	}

	// Create entity
	entity := input.ToEntity()

	// Validate entity
	errs = entity.Validate()
	if errs != nil {
		output.SetErrors(shared.DomainCodeInvalidEntity, errs)
		return
	}

	// Save entity
	err := c.Repository.Save(entity)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, "Internal error")
		return
	}

	// Return result
	output.SetOk()
	return
}

func (c *ProductUseCase) Query(input QueryProductInput) (output shared.Output) {
	// Validate input (fail fast)
	errs := input.Validate()
	if errs != nil {
		output.SetErrors(shared.DomainCodeInvalidInput, errs)
		return
	}

	// Query entities
	entities, err := c.Repository.Query(input.From, input.To)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, "Internal error")
		return
	}

	// Return result
	output.SetOkWithData(entities)
	return
}
