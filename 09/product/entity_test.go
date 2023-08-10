package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_Success(t *testing.T) {
	// When | Arrange
	name := "Product 1"
	price := 0.01

	// Given | Act
	entity, err := NewProduct(name, price)

	// Then | Assert
	assert.Nil(t, err)
	assert.Equal(t, name, entity.Name)
	assert.Equal(t, price, entity.Price)
	assert.NotEmpty(t, entity.ID)
	assert.NotEmpty(t, entity.CreatedAt)
}

func TestCreateProduct_Invalid_Name(t *testing.T) {
	// When | Arrange
	name := ""
	price := 0.01

	// Given | Act
	_, err := NewProduct(name, price)

	// Then | Assert
	assert.NotNil(t, err)
	assert.Equal(t, "name is required", err.Error())
}

func TestCreateProduct_Invalid_Price(t *testing.T) {
	// When | Arrange
	name := "Product 1"
	price := 0.00

	// Given | Act
	_, err := NewProduct(name, price)

	// Then | Assert
	assert.NotNil(t, err)
	assert.Equal(t, "price is required", err.Error())
}

func TestCreateProduct_Invalid_Name_And_Price(t *testing.T) {
	// When | Arrange
	name := ""
	price := 0.00

	// Given | Act
	_, err := NewProduct(name, price)

	// Then | Assert
	assert.NotNil(t, err)
	assert.Equal(t, "name is required, price is required", err.Error())
}
