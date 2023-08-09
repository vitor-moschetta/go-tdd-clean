package product

type ProductUseCase struct {
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

	// save entity to storage

	// return output
	return nil
}
