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
	repository := NewProductRepositoryInMemory()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Nil(t, output)
}
