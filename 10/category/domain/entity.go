package category

import (
	"go-tdd-clean/10/shared"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        string
	Name      string
	CreatedAt string
}

func NewCategory(name string) (Category, error) {
	category := Category{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now().Format(time.RFC3339), // 2021-01-01 00:00:00
	}
	err := category.Validate()
	return category, err
}

func (p *Category) Validate() error {
	err := new(shared.Error)
	if p.Name == "" {
		err.AddError("name is required")
	}
	if p.ID == "" {
		err.AddError("id is required")
	}
	if p.CreatedAt == "" {
		err.AddError("created_at is required")
	}
	if err.Error() != "" {
		return err
	}
	return nil
}
