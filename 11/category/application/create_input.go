package category

import category "go-tdd-clean/11/category/domain"

type CreateCategoryInput struct {
	Name string
}

func (c *CreateCategoryInput) Validate() (errs []string) {
	if c.Name == "" {
		errs = append(errs, "name is required")
	}
	return errs
}

func (c *CreateCategoryInput) ToEntity() category.Category {
	return category.NewCategory(c.Name)
}
