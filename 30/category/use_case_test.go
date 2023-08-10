package category

import (
	"go-tdd-clean/09/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_ValidInput(t *testing.T) {
	// When | Arrange
	input := CreateCategoryInput{
		Name: "Category A",
	}
	repository := NewCategoryRepositoryInMemory()
	useCase := NewCreateCategoryUseCase(repository)

	// Given | Act
	output := useCase.Execute(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, shared.DomainCodeSuccess, output.GetCode())
}
