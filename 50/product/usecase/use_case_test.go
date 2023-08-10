package product

import (
	category "go-tdd-clean/50/category/domain"
	categoryUseCase "go-tdd-clean/50/category/usecase"
	product "go-tdd-clean/50/product/domain"

	"go-tdd-clean/50/shared"
	"go-tdd-clean/50/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductUseCase_Success(t *testing.T) {
	// When | Arrange
	input := CreateProductInput{
		Name:         "Product 1",
		Price:        500,
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
