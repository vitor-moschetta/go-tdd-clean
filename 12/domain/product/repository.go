package product

type IProductRepository interface {
	Save(item Product) error
	QueryMinMaxPrice(min float64, max float64) ([]Product, error)
	Query(func(Product) bool) ([]Product, error)
}
