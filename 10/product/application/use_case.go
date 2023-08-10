package product

import (
	"errors"
	category "go-tdd-clean/10/category/application"
	domainCategory "go-tdd-clean/10/category/domain"
	domain "go-tdd-clean/10/product/domain"
	"go-tdd-clean/10/shared"
	"log"
)

type CreateProductUseCase struct {
	Repository domain.IProductRepository
	shared.Mediator
}

func NewCreateProductUseCase(repository domain.IProductRepository, mediator *shared.Mediator) *CreateProductUseCase {
	return &CreateProductUseCase{
		Repository: repository,
		Mediator:   *mediator,
	}
}

func (c *CreateProductUseCase) Execute(in any) (output shared.Output) {
	// validate input
	input, ok := in.(CreateProductInput)
	if !ok {
		output.SetError(shared.DomainCodeInvalidInput, errors.New("invalid input"))
		return
	}
	err := input.Validate()
	if err != nil {
		output.SetError(shared.DomainCodeInvalidInput, err)
		return
	}

	// check if product already exists
	fn := func(p domain.Product) bool {
		return p.Name == input.Name
	}
	entities, err := c.Repository.Query(fn)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}
	if entities != nil && len(entities) > 0 {
		output.SetError(shared.DomainCodeInternalError, errors.New("product already exists"))
		return
	}

	// check if category exists
	fn = func(p domain.Product) bool {
		return p.CategoryID == input.CategoryID
	}
	entities, err = c.Repository.Query(fn)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}
	if entities == nil || len(entities) == 0 {
		out := c.Mediator.Execute(shared.CreateCategoryUseCase, category.CreateCategoryInput{
			Name: input.CategoryName,
		})
		if out.GetCode() != shared.DomainCodeSuccess {
			output.SetError(shared.DomainCodeInternalError, errors.New("category not found and could not be created"))
			return
		}
		input.CategoryID = out.GetData().(domainCategory.Category).ID
	}

	// create entity
	entity, err := domain.NewProduct(input.Name, input.Price, input.CategoryID)
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

type GetProductByMinMaxPriceUseCase struct {
	Repository domain.IProductRepository
}

func (c *GetProductByMinMaxPriceUseCase) Execute(input GetProductByMinMaxPriceInput) (output shared.Output) {
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
