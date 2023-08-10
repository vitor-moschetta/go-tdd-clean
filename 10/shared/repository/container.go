package repository

import (
	category "go-tdd-clean/10/category/domain"
	product "go-tdd-clean/10/product/domain"
)

type RepositoryContainer struct {
	ProductRepo  product.IProductRepository
	CategoryRepo category.ICategoryRepository
}

func NewRepositoryContainer(
	product product.IProductRepository,
	category category.ICategoryRepository,
) *RepositoryContainer {
	return &RepositoryContainer{
		ProductRepo:  product,
		CategoryRepo: category,
	}
}
