package category

import (
	"go-tdd-clean/10/category"
	"go-tdd-clean/10/shared"
	"go-tdd-clean/10/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCategoryUseCase_Success(t *testing.T) {
	// When | Arrange
	input := category.CreateCategoryInput{
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
