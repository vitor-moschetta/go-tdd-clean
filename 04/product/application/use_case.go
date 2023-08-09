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

func (c *ProductUseCase) Create(input CreateProductInput) error {
	// validate input
	err := input.Validate()
	if err != nil {
		return err
	}

	// create entity
	entity := product.Product{
		Name: input.Name,
	}

	// save entity to storage
	err = c.Repository.Save(entity)
	if err != nil {
		log.Println(err)
		return err
	}

	// return output
	return nil
}
