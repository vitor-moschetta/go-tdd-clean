package product

// ====================================== Command ======================================
type CreateProductInput struct {
	Name  string
	Price float64
}

func NewCreateProductInput(name string, price float64) CreateProductInput {
	return CreateProductInput{
		Name:  name,
		Price: price,
	}
}

func (c *CreateProductInput) Validate() error {
	err := new(Error)
	if c.Name == "" {
		err.AddError("name is required")
	}
	if c.Price <= 0 {
		err.AddError("price is required")
	}
	if err.Error() != "" {
		return err
	}
	return nil
}

// ====================================== Query ======================================
type GetProductByMinMaxPriceInput struct {
	MinPrice float64
	MaxPrice float64
}

func NewGetProductByMinMaxPriceInput(minPrice float64, maxPrice float64) GetProductByMinMaxPriceInput {
	return GetProductByMinMaxPriceInput{
		MinPrice: minPrice,
		MaxPrice: maxPrice,
	}
}

func (c *GetProductByMinMaxPriceInput) Validate() error {
	err := new(Error)
	if c.MinPrice < 0 {
		err.AddError("min is negative")
	}
	if c.MaxPrice < 0 {
		err.AddError("max is negative")
	}
	if c.MinPrice > c.MaxPrice {
		err.AddError("min is greater than max")
	}
	if err.Error() != "" {
		return err
	}
	return nil
}
