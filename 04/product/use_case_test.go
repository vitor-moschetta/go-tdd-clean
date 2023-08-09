package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_Success(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name: "Product 1",
	}
	repository := NewProductRepositoryFake()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Nil(t, output)
}

func TestCreateProduct_InvalidInput(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name: "",
	}
	repository := NewProductRepositoryFake()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, "name is required", output.Error())
}