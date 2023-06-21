package test

import (
	"fmt"
	"strings"
	"testing"

	"go-tdd-clean/08/mock"
	"go-tdd-clean/08/product"
	"go-tdd-clean/08/shared"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_ValidInput(t *testing.T) {
	// When | Arrange
	input := product.CreateProductInput{
		Name:  "Product 1",
		Price: 100,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := product.NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Equal(t, 0, len(output.GetErrors()))
	assert.Equal(t, shared.DomainCodeSuccess, output.GetCode())
}

func TestCreateProduct_InvalidInput_Name(t *testing.T) {
	// When | Arrange
	input := product.CreateProductInput{
		Name:  "",
		Price: 100,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := product.NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Equal(t, 1, len(output.GetErrors()))
	assert.Equal(t, shared.DomainCodeInvalidInput, output.GetCode())
	assert.Equal(t, "name is required", output.GetErrors()[0])
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
	assert.Equal(t, 1, len(output.GetErrors()))
	assert.Equal(t, shared.DomainCodeInvalidInput, output.GetCode())
	assert.Equal(t, "price must be greater than 0", output.GetErrors()[0])
}

func TestCreateProduct_InvalidInput_NameAndPrice(t *testing.T) {
	// When | Arrange
	input := product.CreateProductInput{
		Name:  "",
		Price: 0,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := product.NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)
	fmt.Println(strings.Join(output.GetErrors(), ", "))

	// Then | Assert
	assert.Equal(t, 2, len(output.GetErrors()))
	assert.Equal(t, shared.DomainCodeInvalidInput, output.GetCode())
	assert.Equal(t, "name is required", output.GetErrors()[0])
	assert.Equal(t, "price must be greater than 0", output.GetErrors()[1])
}

func TestCreateProduct_InternalError(t *testing.T) {
	// When | Arrange
	input := product.CreateProductInput{
		Name:  "Product 1",
		Price: 100,
	}
	repository := mock.NewProductRepositoryFake()
	repository.SetError()
	useCase := product.NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Equal(t, 1, len(output.GetErrors()))
	assert.Equal(t, shared.DomainCodeInternalError, output.GetCode())
	assert.Equal(t, "Internal error", output.GetErrors()[0])
}
