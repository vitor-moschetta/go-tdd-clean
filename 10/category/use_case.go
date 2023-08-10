package category

import (
	"errors"
	"go-tdd-clean/10/shared"
	"log"
)

type CreateCategoryUseCase struct {
	repository ICategoryRepository
}

func NewCreateCategoryUseCase(repository ICategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		repository: repository,
	}
}

func (p *CreateCategoryUseCase) Execute(in any) (output shared.Output) {
	input, ok := in.(CreateCategoryInput)
	if !ok {
		output.SetError(shared.DomainCodeInvalidInput, errors.New("invalid category input"))
		return
	}

	// validate input
	err := input.Validate()
	if err != nil {
		output.SetError(shared.DomainCodeInvalidInput, err)
		return
	}

	// create entity
	entity, err := NewCategory(input.Name)
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

type GetCategoryByIDUseCase struct {
	repository ICategoryRepository
}

func NewGetCategoryByIDUseCase(repository ICategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		repository: repository,
	}
}

func (p *GetCategoryByIDUseCase) Execute(input GetCategoryByID) (output shared.Output) {
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
