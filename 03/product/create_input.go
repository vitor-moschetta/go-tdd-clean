package product

type CreateProductInput struct {
	Name string
}

func (c *CreateProductInput) Validate() bool {
	if c.Name == "" {
		return false
	}
	return true
}
