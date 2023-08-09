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

func (c *ProductUseCase) Create(input CreateProductInput) (entity Product, err error) {
	// validate input
	err = input.Validate()
	if err != nil {
		return
	}

	// create entity
	entity, err = NewProduct(input.Name, input.Price)
	if err != nil {
		log.Println(err)
		return
	}

	// save entity to storage
	err = c.Repository.Save(entity)
	if err != nil {
		log.Println(err)
	}

	// return output
	return
}
