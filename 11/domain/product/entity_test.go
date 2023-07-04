package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_Success(t *testing.T) {
	// When | Arrange
	name := "Product 1"
	price := 0.01
	entity := NewProduct(name, price)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 0, len(errs))
	assert.NotNil(t, entity)
	assert.Equal(t, name, entity.Name)
	assert.Equal(t, price, entity.Price)
	assert.NotNil(t, entity.ID)
	assert.NotEmpty(t, entity.ID)
	assert.NotNil(t, entity.CreatedAt)
}

func TestCreateProduct_Invalid_Name(t *testing.T) {
	// When | Arrange
	name := ""
	price := 0.01
	entity := NewProduct(name, price)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "name is required", errs[0])
}

func TestCreateProduct_Invalid_Name2(t *testing.T) {
	// When | Arrange
	var name string
	price := 0.01
	entity := NewProduct(name, price)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "name is required", errs[0])
}

func TestCreateProduct_Invalid_Price(t *testing.T) {
	// When | Arrange
	name := "Product 1"
	price := 0.00
	entity := NewProduct(name, price)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "price must be greater than 0", errs[0])
}

func TestCreateProduct_Invalid_Price2(t *testing.T) {
	// When | Arrange
	name := "Product 1"
	var price float64
	entity := NewProduct(name, price)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "price must be greater than 0", errs[0])
}

func TestCreateProduct_Invalid_Name_And_Price(t *testing.T) {
	// When | Arrange
	name := ""
	price := 0.00
	entity := NewProduct(name, price)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 2, len(errs))
	assert.Equal(t, "name is required", errs[0])
	assert.Equal(t, "price must be greater than 0", errs[1])
}

func TestCreateProduct_Invalid_Name_And_Price2(t *testing.T) {
	// When | Arrange
	var name string
	var price float64
	entity := NewProduct(name, price)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 2, len(errs))
	assert.Equal(t, "name is required", errs[0])
	assert.Equal(t, "price must be greater than 0", errs[1])
}

func TestCreateProduct_Invalid_ID(t *testing.T) {
	// When | Arrange
	name := "Product 1"
	price := 0.01
	entity := NewProduct(name, price)
	entity.ID = ""

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "id is required", errs[0])
}

func TestCreateProduct_Invalid_CreatedAt(t *testing.T) {
	// When | Arrange
	name := "Product 1"
	price := 0.01
	entity := NewProduct(name, price)
	entity.CreatedAt = ""

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "created_at is required", errs[0])
}
