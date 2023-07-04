package product

type QueryProductFromToDate struct {
	From string
	To   string
}

func (c *QueryProductFromToDate) Validate() (errs []string) {
	if c.From == "" {
		errs = append(errs, "from is empty")
	}
	if c.To == "" {
		errs = append(errs, "to is empty")
	}
	return
}
