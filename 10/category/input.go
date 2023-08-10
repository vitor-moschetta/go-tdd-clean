package category

import "go-tdd-clean/10/shared"

// ====================================== Command ======================================
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

type GetCategoryByID struct {
	ID string
}

func (c *GetCategoryByID) Validate() error {
	err := new(shared.Error)
	if c.ID == "" {
		err.AddError("id is required")
	}
	if err.Error() != "" {
		return err
	}
	return nil
}
