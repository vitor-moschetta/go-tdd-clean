package product

import (
	"log"
)

type ProductUseCase struct {
	Repository IProductRepository
}

func NewProductUseCase(repository IProductRepository) *ProductUseCase {
	return &ProductUseCase{
		Repository: repository,
	}
}

func (c *ProductUseCase) Execute(input CreateProductInput) error {
	// validate input
	err := input.Validate()
	if err != nil {
		return err
	}

	// create entity
	entity, err := NewProduct(input.Name, input.Price)
	if err != nil {
		log.Println(err)
		return err
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
