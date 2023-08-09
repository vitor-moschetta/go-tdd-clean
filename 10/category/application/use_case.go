package category

import (
	category "go-tdd-clean/10/category/domain"
	"go-tdd-clean/10/shared"
	"log"
)

type CategoryUseCase struct {
	repository category.ICategoryRepository
}

func NewCategoryUseCase(repository category.ICategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		repository: repository,
	}
}

func (p *CategoryUseCase) CreateCategory(input CreateCategoryInput) (output shared.Output) {
	// validate input
	err := input.Validate()
	if err != nil {
		output.SetError(shared.DomainCodeInvalidInput, err)
		return
	}

	// create entity
	entity, err := category.NewCategory(input.Name)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}

	// save entity to storage
	err = p.repository.Save(entity)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}

	// return ok
	output.SetOk(entity)
	return
}

func (p *CategoryUseCase) GetCategoryByID(input GetCategoryByID) (output shared.Output) {
	// validate input (fail fast)
	err := input.Validate()
	if err != nil {
		output.SetError(shared.DomainCodeInvalidInput, err)
		return
	}

	// query entities
	entity, err := p.repository.GetByID(input.ID)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}

	// return ok
	output.SetOk(entity)
	return
}
