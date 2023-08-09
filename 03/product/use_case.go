package product

import "log"

type ProductUseCase struct {
	Repository IProductRepository
}

func NewProductUseCase() *ProductUseCase {
	return &ProductUseCase{}
}

func (c *ProductUseCase) Create(input CreateProductInput) error {
	// validate input
	err := input.Validate()
	if err != nil {
		return err
	}

	// create entity
	entity := Product{
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
