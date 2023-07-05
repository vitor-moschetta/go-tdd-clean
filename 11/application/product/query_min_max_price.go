package product

type QueryProductMinMaxPrice struct {
	MinPrice float64
	MaxPrice float64
}

func (c *QueryProductMinMaxPrice) Validate() (errs []string) {
	if c.MinPrice < 0 {
		errs = append(errs, "min is negative")
	}
	if c.MaxPrice < 0 {
		errs = append(errs, "max is negative")
	}
	if c.MinPrice > c.MaxPrice {
		errs = append(errs, "min is greater than max")
	}
	return
}
