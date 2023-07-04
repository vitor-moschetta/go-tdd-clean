package product

type IProductRepository interface {
	Save(item Product) error
	Query(from string, to string) ([]Product, error)
}
