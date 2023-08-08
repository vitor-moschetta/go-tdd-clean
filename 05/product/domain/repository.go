package product

type IProductRepository interface {
	Save(item Product) error
}
