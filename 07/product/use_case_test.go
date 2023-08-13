package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_ValidInput(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "Product 1",
		Price: 100,
	}
	repository := NewInMemoryProductRepository()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Execute(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, DomainCodeSuccess, output.GetCode())
	assert.Equal(t, input.Name, output.GetData().(Product).Name)
	assert.Equal(t, input.Price, output.GetData().(Product).Price)
}

func TestCreateProduct_InvalidInput(t *testing.T) {
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
	assert.Equal(t, DomainCodeInvalidInput, output.GetCode())
	assert.Equal(t, "name is required, price is required", output.GetError())
}
