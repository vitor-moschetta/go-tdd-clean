package test

import (
	"testing"

	"go-tdd-clean/07/mock"
	"go-tdd-clean/07/product"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_ValidInput(t *testing.T) {
	// When | Arrange
	input := product.CreateProductInput{
		Name:  "Product 1",
		Price: 9,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := product.NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Equal(t, 0, len(output))
}

func TestCreateProduct_InvalidInput(t *testing.T) {
	// When | Arrange
	input := product.CreateProductInput{
		Name:  "",
		Price: 0,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := product.NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Equal(t, 2, len(output))
}

func TestCreateProduct_InvalidInput_Name(t *testing.T) {
	// When | Arrange
	input := product.CreateProductInput{
		Name:  "",
		Price: 9,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := product.NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Equal(t, 1, len(output))
	assert.Equal(t, "name is required", output[0].Error())
}

func TestCreateProduct_InvalidInput_Price(t *testing.T) {
	// When | Arrange
	input := product.CreateProductInput{
		Name:  "Product 1",
		Price: 0,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := product.NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Equal(t, 1, len(output))
	assert.Equal(t, "price must be greater than 0", output[0].Error())
}
