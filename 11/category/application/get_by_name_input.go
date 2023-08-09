package category

type GetCategoryByName struct {
	Name string
}

func (c *GetCategoryByName) Validate() (errs []string) {
	if c.Name == "" {
		errs = append(errs, "name is required")
	}
	return errs
}
