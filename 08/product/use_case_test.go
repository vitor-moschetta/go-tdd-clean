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
	repository := NewProductRepositoryFake()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

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
	repository := NewProductRepositoryFake()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, DomainCodeInvalidInput, output.GetCode())
	assert.Equal(t, "name is required, price is required", output.GetError())
}

func TestGetByMinMaxPrice_ValidInput(t *testing.T) {
	// When | Arrange
	input := GetProductByMinMaxPriceInput{
		MinPrice: 100,
		MaxPrice: 200,
	}
	repository := NewProductRepositoryFake()
	repository.Seed()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.GetByMinMaxPrice(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, DomainCodeSuccess, output.GetCode())
	assert.Equal(t, 2, len(output.GetData().([]Product)))
}

func TestGetByMinMaxPrice_ValidInput2(t *testing.T) {
	// When | Arrange
	input := GetProductByMinMaxPriceInput{
		MinPrice: 0,
		MaxPrice: 100,
	}
	repository := NewProductRepositoryFake()
	repository.Seed()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.GetByMinMaxPrice(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, DomainCodeSuccess, output.GetCode())
	assert.Equal(t, 1, len(output.GetData().([]Product)))
}

func TestGetByMinMaxPrice_InvalidInput(t *testing.T) {
	// When | Arrange
	input := GetProductByMinMaxPriceInput{
		MinPrice: 200,
		MaxPrice: 100,
	}
	repository := NewProductRepositoryFake()
	repository.Seed()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.GetByMinMaxPrice(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, DomainCodeInvalidInput, output.GetCode())
	assert.Equal(t, "min is greater than max", output.GetError())
}
