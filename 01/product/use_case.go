package product

type ProductUseCase struct {
}

func NewProductUseCase() *ProductUseCase {
	return &ProductUseCase{}
}

func (c *ProductUseCase) Create(input CreateProductInput) (output error) {
	return nil
}
