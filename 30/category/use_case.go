package category

import (
	"go-tdd-clean/09/shared"
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

func (p *CreateCategoryUseCase) Execute(input CreateCategoryInput) (output shared.Output) {
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
