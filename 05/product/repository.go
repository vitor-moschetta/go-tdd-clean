package product

type IProductRepository interface {
	Save(item Product) error
	GetByName(name string) (Product, error)
}
