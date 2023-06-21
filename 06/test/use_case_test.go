package test

import (
	"testing"

	"go-tdd-clean/06/mock"
	"go-tdd-clean/06/product"

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
	assert.True(t, output)
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
	assert.False(t, output)
}
