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
	errrs := input.Validate()

	// Then | Assert
	assert.Equal(t, 0, len(errrs))
}

func TestCreateInput_Invalid_Name(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "",
		Price: 100,
	}

	// Given | Act
	errrs := input.Validate()

	// Then | Assert
	assert.Equal(t, 1, len(errrs))
	assert.Equal(t, "name is required", errrs[0])
}

func TestCreateInput_Invalid_Price(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "Product 1",
		Price: 0,
	}

	// Given | Act
	errrs := input.Validate()

	// Then | Assert
	assert.Equal(t, 1, len(errrs))
	assert.Equal(t, "price must be greater than 0", errrs[0])
}

func TestCreateInput_Invalid_Name_And_Price(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "",
		Price: 0,
	}

	// Given | Act
	errrs := input.Validate()

	// Then | Assert
	assert.Equal(t, 2, len(errrs))
	assert.Equal(t, "name is required", errrs[0])
	assert.Equal(t, "price must be greater than 0", errrs[1])
}
