package usecase

import (
	category "go-tdd-clean/50/category/domain"
	"go-tdd-clean/50/shared"
	"go-tdd-clean/50/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCategoryUseCase_Success(t *testing.T) {
	// When | Arrange
	input := CreateCategoryInput{
		Name: "Category A",
	}
	categoryRepo := category.NewCategoryRepositoryInMemory()
	repoContainer := repository.NewRepositoryContainer(nil, categoryRepo)
	useCase := NewCreateCategoryUseCase(repoContainer)

	// Given | Act
	output := useCase.Execute(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, shared.DomainCodeSuccess, output.GetCode())
}
