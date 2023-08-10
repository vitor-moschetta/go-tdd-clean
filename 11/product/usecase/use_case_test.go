package product

import (
	category "go-tdd-clean/11/category/domain"
	categoryUseCase "go-tdd-clean/11/category/usecase"
	product "go-tdd-clean/11/product/domain"

	"go-tdd-clean/11/shared"
	"go-tdd-clean/11/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductUseCase_Success(t *testing.T) {
	// When | Arrange
	input := product.CreateProductInput{
		Name:         "Product 1",
		Price:        110,
		CategoryID:   "1",
		CategoryName: "Category A",
	}
	productRepo := product.NewProductRepositoryInMemory()
	categoryRepo := category.NewCategoryRepositoryInMemory()
	repoContainer := repository.NewRepositoryContainer(productRepo, categoryRepo)
	mediator := shared.NewMediator()
	createProductUseCase := NewCreateProductUseCase(repoContainer, mediator)
	mediator.RegisterUseCase(shared.CreateProductUseCaseKey, createProductUseCase)
	createCategoryUseCase := categoryUseCase.NewCreateCategoryUseCase(repoContainer)
	mediator.RegisterUseCase(shared.CreateCategoryUseCaseKey, createCategoryUseCase)

	// Given | Act
	output := createProductUseCase.Execute(input)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, shared.DomainCodeSuccess, output.GetCode())
	assert.Equal(t, input.Name, output.GetData().(product.Product).Name)
	assert.Equal(t, input.Price, output.GetData().(product.Product).Price)
	category, err := categoryRepo.GetByID(output.GetData().(product.Product).CategoryID)
	assert.Nil(t, err)
	assert.NotNil(t, category)
}
