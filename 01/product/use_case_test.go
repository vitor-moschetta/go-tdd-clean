package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name: "Product 1",
	}
	useCase := NewProductUseCase()

	// Given | Act
	output := useCase.Execute(input)

	// Then | Assert
	assert.Nil(t, output)
}
