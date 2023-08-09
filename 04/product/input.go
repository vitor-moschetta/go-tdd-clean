package product

import "errors"

type CreateProductInput struct {
	Name string
}

func (c *CreateProductInput) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
