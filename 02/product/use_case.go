package product

type ProductUseCase struct {
}

func NewProductUseCase() *ProductUseCase {
	return &ProductUseCase{}
}

func (c *ProductUseCase) Create(input CreateProductInput) (output error) {
	// Validate input
	err := input.Validate()
	if err != nil {
		return err
	}

	// Create entity

	// Save entity to storage

	// Return result
	return nil
}
