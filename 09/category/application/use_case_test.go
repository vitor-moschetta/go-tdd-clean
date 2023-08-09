package category

import (
	mock "go-tdd-clean/09/category/infrastructure"
	"go-tdd-clean/09/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_ValidInput(t *testing.T) {
	// When | Arrange
	input := CreateCategoryInput{
		Name: "Category A",
	}
	repository := mock.NewCategoryRepositoryFake()
	useCase := NewCategoryUseCase(repository)

	// Given | Act
	output := useCase.CreateCategory(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, shared.DomainCodeSuccess, output.GetCode())
}
