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
	repository := NewProductRepositoryFake()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.Create(input)

	// Then | Assert
	assert.Nil(t, output)
}

// Problema: como sei o tipo de erro que o use case retornou?
// Sem saber o tipo de erro, como vou mapear o código HTTP correto (exemplo caso esteja usando API REST na minha infraestrutura)?
// E se for necessário retornar mais informações de domínio?
