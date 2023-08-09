package category

import (
	category "go-tdd-clean/11/category/domain"

	"go-tdd-clean/11/shared"
	"log"
)

type ProductUseCase struct {
	Repository category.ICategoryRepository
}

func NewProductUseCase(repository category.ICategoryRepository) *ProductUseCase {
	return &ProductUseCase{
		Repository: repository,
	}
}

func (c *ProductUseCase) Create(input CreateCategoryInput) (output shared.Output) {
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

func (c *ProductUseCase) GetByName(input GetCategoryByName) (output shared.Output) {
	// Validate input (fail fast)
	errs := input.Validate()
	if errs != nil {
		output.SetErrors(shared.DomainCodeInvalidInput, errs)
		return
	}

	entity, err := c.Repository.GetByName(input.Name)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, "Internal error")
		return
	}

	// Return result
	output.SetOkWithData(entity)
	return
}
