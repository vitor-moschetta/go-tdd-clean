package product

import (
	domain "go-tdd-clean/09/product/domain"
	"go-tdd-clean/09/shared"
	"log"
)

type ProductUseCase struct {
	Repository domain.IProductRepository
}

func NewProductUseCase(repository domain.IProductRepository) *ProductUseCase {
	return &ProductUseCase{
		Repository: repository,
	}
}

func (c *ProductUseCase) Create(input CreateProductInput) (output shared.Output) {
	// validate input
	err := input.Validate()
	if err != nil {
		output.SetError(shared.DomainCodeInvalidInput, err)
		return
	}

	// create entity
	entity, err := domain.NewProduct(input.Name, input.Price)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}

	// save entity to storage
	err = c.Repository.Save(entity)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}

	// return ok
	output.SetOk(entity)
	return
}

func (c *ProductUseCase) GetByMinMaxPrice(input GetProductByMinMaxPriceInput) (output shared.Output) {
	// validate input (fail fast)
	err := input.Validate()
	if err != nil {
		output.SetError(shared.DomainCodeInvalidInput, err)
		return
	}

	// query entities
	fn := func(p domain.Product) bool {
		return p.Price >= input.MinPrice && p.Price <= input.MaxPrice
	}

	entities, err := c.Repository.Query(fn)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}

	// return ok
	output.SetOk(entities)
	return
}
