package product

type IProductRepository interface {
	Save(item Product) error
	GetAll() ([]Product, error)
	Query(fn func(Product) bool) ([]Product, error)
}
