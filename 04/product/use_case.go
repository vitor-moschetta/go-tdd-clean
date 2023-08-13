package product

import (
	"errors"
	"log"

	"github.com/google/uuid"
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

	// verify if product already exists
	entity, err := c.Repository.GetByName(input.Name)
	if err != nil {
		log.Println(err)
		return err
	}
	if entity != nil {
		return errors.New("product already exists")
	}

	// create entity
	entity = &Product{
		ID:   uuid.New().String(),
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
