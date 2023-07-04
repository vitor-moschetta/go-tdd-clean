package product

type QueryProductMinMaxPrice struct {
	Min float64
	Max float64
}

func (c *QueryProductMinMaxPrice) Validate() (errs []string) {
	if c.Min < 0 {
		errs = append(errs, "min is negative")
	}
	if c.Max < 0 {
		errs = append(errs, "max is negative")
	}
	if c.Min > c.Max {
		errs = append(errs, "min is greater than max")
	}
	return
}
