package product

import (
	"go-tdd-clean/10/category"
	"go-tdd-clean/10/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductUseCase_Success(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:         "Product 1",
		Price:        100,
		CategoryID:   "1",
		CategoryName: "Category A",
	}
	productRepository := NewProductRepositoryInMemory()
	categoryRepository := category.NewCategoryRepositoryInMemory()
	mediator := shared.NewMediator()
	createProductUseCase := NewCreateProductUseCase(productRepository, mediator)
	mediator.RegisterUseCase(shared.CreateProductUseCaseKey, createProductUseCase)
	createCategoryUseCase := category.NewCreateCategoryUseCase(categoryRepository)
	mediator.RegisterUseCase(shared.CreateCategoryUseCaseKey, createCategoryUseCase)

	// Given | Act
	output := createProductUseCase.Execute(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, shared.DomainCodeSuccess, output.GetCode())
	assert.Equal(t, input.Name, output.GetData().(Product).Name)
	assert.Equal(t, input.Price, output.GetData().(Product).Price)
	category, err := categoryRepository.GetByID(output.GetData().(Product).CategoryID)
	assert.Nil(t, err)
	assert.NotNil(t, category)
}
