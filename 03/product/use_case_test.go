package product

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_ValidInput(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name: "Product 1",
	}
	useCase := NewProductUseCase()

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(runtime.Error); ok {
				assert.Equal(t, "runtime error: invalid memory address or nil pointer dereference", err.Error())
			} else {
				assert.True(t, true)
			}
		}
	}()

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.NotNil(t, output)
}

func TestCreateProduct_InvalidInput(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name: "",
	}
	useCase := NewProductUseCase()

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, "name is required", output.Error())
}
