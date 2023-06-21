package product

import "log"

type ProductUseCase struct {
	Repository IProductRepository
}

func NewProductUseCase(repository IProductRepository) *ProductUseCase {
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
	entity := NewProduct(input.Name, input.Price)

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
