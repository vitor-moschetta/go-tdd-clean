package product

import (
	"testing"
	"time"

	"go-tdd-clean/11/domain/product"
	"go-tdd-clean/11/mock"
	"go-tdd-clean/11/shared"

	"github.com/stretchr/testify/assert"
)

func TestUseCase_CreateProduct_Success(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "Product 1",
		Price: 0.01,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Equal(t, 0, len(output.GetErrors()))
	assert.Equal(t, shared.DomainCodeSuccess, output.GetCode())
}

func TestUseCase_CreateProduct_InvalidInput(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "Product 1",
		Price: 0,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Equal(t, 1, len(output.GetErrors()))
	assert.Equal(t, shared.DomainCodeInvalidInput, output.GetCode())
}

func TestUseCase_CreateProduct_InvalidInput2(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "",
		Price: 0,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Equal(t, 2, len(output.GetErrors()))
	assert.Equal(t, shared.DomainCodeInvalidInput, output.GetCode())
}

func TestCreateProduct_InternalError(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "Product 1",
		Price: 100,
	}
	repository := mock.NewProductRepositoryFake()
	repository.SetError()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Equal(t, 1, len(output.GetErrors()))
	assert.Equal(t, shared.DomainCodeInternalError, output.GetCode())
	assert.Equal(t, "Internal error", output.GetErrors()[0])
}

func TestUseCase_QueryProduct_Success(t *testing.T) {
	// When | Arrange
	input := QueryProductInput{
		From: string(time.Now().AddDate(0, 0, -1).Format("2006-01-02")),
		To:   string(time.Now().AddDate(0, 0, 1).Format("2006-01-02")),
	}
	repository := mock.NewProductRepositoryFake()
	repository.Seed()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Query(input)

	// Then | Assert
	assert.Equal(t, 0, len(output.GetErrors()))
	assert.Equal(t, shared.DomainCodeSuccess, output.GetCode())
	assert.Equal(t, 2, len(output.GetData().([]product.Product)))
}
