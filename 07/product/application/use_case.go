package product

import (
	product "go-tdd-clean/07/product/domain"
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

func (c *ProductUseCase) Create(input CreateProductInput) (errs []error) {
	// Validate input (fail fast)
	errs = input.Validate()
	if len(errs) > 0 {
		return errs
	}

	// Create entity
	entity := product.NewProduct(input.Name, input.Price)

	// Validate entity
	errs = entity.Validate()
	if len(errs) > 0 {
		return errs
	}

	// Save entity
	err := c.Repository.Save(entity)
	if err != nil {
		log.Println(err)
		return []error{err}
	}

	// Return result
	return errs
}
