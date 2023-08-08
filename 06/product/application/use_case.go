package product

import (
	product "go-tdd-clean/06/product/domain"
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
	// Validate input (fail fast)
	if !input.Validate() {
		return false
	}

	// Create entity
	entity := product.NewProduct(input.Name, input.Price)

	// Validate entity
	if !entity.Validate() {
		return false
	}

	// Save entity
	err := c.Repository.Save(entity)
	if err != nil {
		log.Println(err)
		return false
	}

	// Return result
	return true
}
