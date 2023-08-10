package usecase

import (
	"errors"

	category "go-tdd-clean/50/category/domain"
	"go-tdd-clean/50/shared"
	"go-tdd-clean/50/shared/repository"
	"log"
)

type CreateCategoryUseCase struct {
	repository.RepositoryContainer
}

func NewCreateCategoryUseCase(repoContainer *repository.RepositoryContainer) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		RepositoryContainer: *repoContainer,
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
	entity, err := category.NewCategory(input.Name)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}

	// save entity to storage
	err = p.CategoryRepo.Save(entity)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}

	// return ok
	output.SetOk(entity)
	return
}
