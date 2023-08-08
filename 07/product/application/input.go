package product

import (
	"errors"
)

type CreateProductInput struct {
	Name  string
	Price float64
}

func (c *CreateProductInput) Validate() (errs []error) {
	if c.Name == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if c.Price <= 0 {
		errs = append(errs, errors.New("price must be greater than 0"))
	}
	return errs
}
