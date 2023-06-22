package product

import (
	"fmt"
	"strings"
	"testing"

	"go-tdd-clean/10/mock"
	"go-tdd-clean/10/shared"

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
		Name:  "",
		Price: 0,
	}
	repository := mock.NewProductRepositoryFake()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)
	fmt.Println(strings.Join(output.GetErrors(), ", "))

	// Then | Assert
	assert.Equal(t, 2, len(output.GetErrors()))
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
