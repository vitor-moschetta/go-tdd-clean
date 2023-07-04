package product

type QueryProductInput struct {
	From string
	To   string
}

func (c *QueryProductInput) Validate() (errs []string) {
	if c.From == "" {
		errs = append(errs, "from is empty")
	}
	if c.To == "" {
		errs = append(errs, "to is empty")
	}
	return
}
