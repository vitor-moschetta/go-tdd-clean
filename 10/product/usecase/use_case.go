package product

import (
	"errors"
	category "go-tdd-clean/10/category/domain"
	product "go-tdd-clean/10/product/domain"
	"go-tdd-clean/10/shared"
	"go-tdd-clean/10/shared/repository"
	"log"
)

type CreateProductUseCase struct {
	repository.RepositoryContainer
	shared.Mediator
}

func NewCreateProductUseCase(repository *repository.RepositoryContainer, mediator *shared.Mediator) *CreateProductUseCase {
	return &CreateProductUseCase{
		RepositoryContainer: *repository,
		Mediator:            *mediator,
	}
}

func (c *CreateProductUseCase) Execute(in any) (output shared.Output) {
	// validate input
	input, ok := in.(product.CreateProductInput)
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
	fn := func(p product.Product) bool {
		return p.Name == input.Name
	}
	entities, err := c.ProductRepo.Query(fn)
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
	categoryEntity, err := c.CategoryRepo.GetByID(input.CategoryID)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}
	if categoryEntity.ID == "" {
		out := c.Mediator.Execute(shared.CreateCategoryUseCaseKey, category.CreateCategoryInput{
			Name: input.CategoryName,
		})
		if out.GetCode() != shared.DomainCodeSuccess {
			output.SetError(shared.DomainCodeInternalError, errors.New("category not found and could not be created"))
			return
		}
		input.CategoryID = out.GetData().(category.Category).ID
	}

	// create entity
	entity, err := product.NewProduct(input.Name, input.Price, input.CategoryID)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}

	// save entity to storage
	err = c.ProductRepo.Save(entity)
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
	repository.RepositoryContainer
}

func NewGetProductByMinMaxPriceUseCase(repoContainer *repository.RepositoryContainer) *GetProductByMinMaxPriceUseCase {
	return &GetProductByMinMaxPriceUseCase{
		RepositoryContainer: *repoContainer,
	}
}

func (c *GetProductByMinMaxPriceUseCase) Execute(in any) (output shared.Output) {
	input, ok := in.(product.GetProductByMinMaxPriceInput)
	if !ok {
		output.SetError(shared.DomainCodeInvalidInput, errors.New("invalid input"))
		return
	}

	// validate input (fail fast)
	err := input.Validate()
	if err != nil {
		output.SetError(shared.DomainCodeInvalidInput, err)
		return
	}

	// query entities
	fn := func(p product.Product) bool {
		return p.Price >= input.MinPrice && p.Price <= input.MaxPrice
	}

	entities, err := c.ProductRepo.Query(fn)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}

	// return ok
	output.SetOk(entities)
	return
}
