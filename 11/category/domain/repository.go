package category

type ICategoryRepository interface {
	Save(item Category) error
	GetByName(name string) (Category, error)
}
