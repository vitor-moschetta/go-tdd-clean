package product

import (
	category "go-tdd-clean/10/category/application"
	mockCat "go-tdd-clean/10/category/infrastructure"
	domain "go-tdd-clean/10/product/domain"
	mockProd "go-tdd-clean/10/product/infrastructure"

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
	productRepository := mockProd.NewProductRepositoryFake()
	categoryRepository := mockCat.NewCategoryRepositoryFake()
	mediator := shared.NewMediator()
	createProductUseCase := NewCreateProductUseCase(productRepository, mediator)
	mediator.RegisterUseCase(shared.CreateCategoryUseCase, createProductUseCase)
	createCategoryUseCase := category.NewCreateCategoryUseCase(categoryRepository)
	mediator.RegisterUseCase(shared.CreateCategoryUseCase, createCategoryUseCase)

	// Given | Act
	output := createProductUseCase.Execute(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, shared.DomainCodeSuccess, output.GetCode())
	assert.Equal(t, input.Name, output.GetData().(domain.Product).Name)
	assert.Equal(t, input.Price, output.GetData().(domain.Product).Price)
	category, err := categoryRepository.GetByID(output.GetData().(domain.Product).CategoryID)
	assert.Nil(t, err)
	assert.NotNil(t, category)
}
