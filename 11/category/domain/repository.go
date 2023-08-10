package category

type ICategoryRepository interface {
	Save(item Category) error
	GetByID(id string) (Category, error)
}
