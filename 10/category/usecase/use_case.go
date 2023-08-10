package category

import (
	"errors"
	"go-tdd-clean/10/category"
	"go-tdd-clean/10/shared"
	"go-tdd-clean/10/shared/repository"
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
	input, ok := in.(category.CreateCategoryInput)
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

type GetCategoryByIDUseCase struct {
	repository.RepositoryContainer
}

func NewGetCategoryByIDUseCase(repoContainer repository.RepositoryContainer) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		RepositoryContainer: repoContainer,
	}
}

func (p *GetCategoryByIDUseCase) Execute(input category.GetCategoryByID) (output shared.Output) {
	// validate input (fail fast)
	err := input.Validate()
	if err != nil {
		output.SetError(shared.DomainCodeInvalidInput, err)
		return
	}

	// query entities
	entity, err := p.CategoryRepo.GetByID(input.ID)
	if err != nil {
		log.Println(err)
		output.SetError(shared.DomainCodeInternalError, err)
		return
	}

	// return ok
	output.SetOk(entity)
	return
}
