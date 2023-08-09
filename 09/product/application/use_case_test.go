package product

import (
	domain "go-tdd-clean/09/product/domain"
	mock "go-tdd-clean/09/product/infrastructure"

	"go-tdd-clean/09/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_ValidInput(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "Product 1",
		Price: 100,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, shared.DomainCodeSuccess, output.GetCode())
	assert.Equal(t, input.Name, output.GetData().(domain.Product).Name)
	assert.Equal(t, input.Price, output.GetData().(domain.Product).Price)
}

func TestCreateProduct_InvalidInput(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "",
		Price: -1,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, shared.DomainCodeInvalidInput, output.GetCode())
	assert.Equal(t, "name is required, price is required", output.GetError())
}

func TestGetByMinMaxPrice_ValidInput(t *testing.T) {
	// When | Arrange
	input := GetProductByMinMaxPriceInput{
		MinPrice: 100,
		MaxPrice: 200,
	}
	repository := mock.NewProductRepositoryFake()
	repository.Seed()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.GetByMinMaxPrice(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, shared.DomainCodeSuccess, output.GetCode())
	assert.Equal(t, 2, len(output.GetData().([]domain.Product)))
}

func TestGetByMinMaxPrice_ValidInput2(t *testing.T) {
	// When | Arrange
	input := GetProductByMinMaxPriceInput{
		MinPrice: 0,
		MaxPrice: 100,
	}
	repository := mock.NewProductRepositoryFake()
	repository.Seed()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.GetByMinMaxPrice(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, shared.DomainCodeSuccess, output.GetCode())
	assert.Equal(t, 1, len(output.GetData().([]domain.Product)))
}

func TestGetByMinMaxPrice_InvalidInput(t *testing.T) {
	// When | Arrange
	input := GetProductByMinMaxPriceInput{
		MinPrice: 200,
		MaxPrice: 100,
	}
	repository := mock.NewProductRepositoryFake()
	repository.Seed()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.GetByMinMaxPrice(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, shared.DomainCodeInvalidInput, output.GetCode())
	assert.Equal(t, "min is greater than max", output.GetError())
}
