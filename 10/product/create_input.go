package product

type CreateProductInput struct {
	Name  string
	Price float64
}

func (c *CreateProductInput) Validate() (errs []string) {
	if c.Name == "" {
		errs = append(errs, "name is required")
	}
	if c.Price <= 0 {
		errs = append(errs, "price must be greater than 0")
	}
	return errs
}

func (c *CreateProductInput) ToEntity() Product {
	return NewProduct(c.Name, c.Price)
}
