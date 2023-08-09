package product

import "errors"

type CreateProductInput struct {
	Name  string
	Price float64
}

func (c *CreateProductInput) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Price <= 0 {
		return errors.New("price is required")
	}
	return nil
}
