package category

import "go-tdd-clean/09/shared"

type CreateCategoryInput struct {
	Name string
}

func (c *CreateCategoryInput) Validate() error {
	err := new(shared.Error)
	if c.Name == "" {
		err.AddError("name is required")
	}
	if err.Error() != "" {
		return err
	}
	return nil
}
