package repository

import (
	category "go-tdd-clean/10/category/repository"
	product "go-tdd-clean/10/product/repository"
)

type RepositoryContainer struct {
	Product  product.IProductRepository
	Category category.ICategoryRepository
}

func NewRepositoryContainer(product product.IProductRepository, category category.ICategoryRepository) *RepositoryContainer {
	return &RepositoryContainer{
		Product:  product,
		Category: category,
	}
}
