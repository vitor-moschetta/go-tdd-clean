package product

type ProductUseCase struct {
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

	// Save entity to storage

	// Return result
	return true
}
