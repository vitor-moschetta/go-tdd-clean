package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_Success(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "Product 1",
		Price: 1000,
	}
	repository := NewInMemoryProductRepository()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Execute(input)

	// Then | Assert
	assert.Nil(t, output)
}

func TestCreateProduct_InvalidInput(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "",
		Price: 1000,
	}
	repository := NewInMemoryProductRepository()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Execute(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, "name is required", output.Error())
}

func TestCreateProduct_InvalidInput2(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "",
		Price: -1,
	}
	repository := NewInMemoryProductRepository()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Execute(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, "name is required, price is required", output.Error())
}
