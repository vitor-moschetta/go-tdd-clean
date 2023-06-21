package product

type CreateProductInput struct {
	Name  string
	Price float64
}

func (c *CreateProductInput) Validate() bool {
	if c.Name == "" {
		return false
	}
	if c.Price <= 0 {
		return false
	}
	return true
}
