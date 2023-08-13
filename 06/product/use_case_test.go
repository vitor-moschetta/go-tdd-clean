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
	repository := NewInMemoryProductRepository()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output, err := useCase.Execute(input)

	// Then | Assert
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotNil(t, output.ID)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Price, output.Price)
}

func TestCreateProduct_InvalidInput(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:  "",
		Price: -1,
	}
	repository := NewInMemoryProductRepository()
	useCase := NewProductUseCase(repository)

	// Given | Act
	_, err := useCase.Execute(input)

	// Then | Assert
	assert.NotNil(t, err)
	assert.Equal(t, "name is required, price is required", err.Error())
}

// Problemas:
// Como sei o tipo de erro que o domínio está retornando?
// Sem saber o tipo de erro, como vou mapear o código HTTP correto (exemplo caso esteja usando API REST na minha infraestrutura)?
// Não seria melhor retornar os erros em uma lista de erros?
