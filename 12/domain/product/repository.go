package product

type IProductRepository interface {
	Save(item Product) error
	Query(func(Product) bool) ([]Product, error)
}
