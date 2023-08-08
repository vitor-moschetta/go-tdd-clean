package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInput(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "Product 1",
		Price: 100,
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 0, len(errs))
}

func TestCreateInput_Invalid_Name(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "",
		Price: 100,
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "name is required", errs[0])
}

func TestCreateInput_Invalid_Price(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "Product 1",
		Price: 0,
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "price must be greater than 0", errs[0])
}

func TestCreateInput_Invalid_Name_And_Price(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "",
		Price: 0,
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 2, len(errs))
	assert.Equal(t, "name is required", errs[0])
	assert.Equal(t, "price must be greater than 0", errs[1])
}
