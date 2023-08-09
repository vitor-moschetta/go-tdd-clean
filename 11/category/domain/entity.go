package category

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        string
	Name      string
	CreatedAt string
}

func NewCategory(name string) Category {
	return Category{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now().Format(time.RFC3339), // 2021-01-01 00:00:00
	}
}

func (p *Category) Validate() (errs []string) {
	if p.Name == "" {
		errs = append(errs, "name is required")
	}
	if p.ID == "" {
		errs = append(errs, "id is required")
	}
	if p.CreatedAt == "" {
		errs = append(errs, "created_at is required")
	}
	return errs
}
