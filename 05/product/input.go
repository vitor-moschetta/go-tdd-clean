package product

type CreateProductInput struct {
	Name  string
	Price float64
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
