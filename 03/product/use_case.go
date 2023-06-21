package product

import "log"

type ProductUseCase struct {
	Repository IProductRepository
}

func NewProductUseCase() *ProductUseCase {
	return &ProductUseCase{}
}

func (c *ProductUseCase) Create(input CreateProductInput) bool {
	// Validate input
	if !input.Validate() {
		return false
	}

	// Create entity
	entity := Product{
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
