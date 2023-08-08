package product

import (
	product "go-tdd-clean/04/product/domain"
	"log"
)

type ProductUseCase struct {
	Repository product.IProductRepository
}

func NewProductUseCase(repository product.IProductRepository) *ProductUseCase {
	return &ProductUseCase{
		Repository: repository,
	}
}

func (c *ProductUseCase) Create(input CreateProductInput) bool {
	// Validate input
	if !input.Validate() {
		return false
	}

	// Create entity
	entity := product.Product{
		Name: input.Name,
	}

	// Save entity to storage
	err := c.Repository.Save(entity)
	if err != nil {
		log.Println(err)
		return false
	}

	// Return result
	return true
}
