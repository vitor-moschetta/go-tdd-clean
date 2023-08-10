package shared

type UseCase interface {
	Execute(input any) Output
}

type UseCaseKey string

// register use cases
const (
	// category
	CreateCategoryUseCaseKey  UseCaseKey = "CreateCategoryUseCase"
	GetCategoryByIDUseCaseKey UseCaseKey = "GetCategoryByIDUseCase"

	// product
	CreateProductUseCaseKey           UseCaseKey = "CreateProductUseCase"
	GetProductByMinMaxPriceUseCaseKey UseCaseKey = "GetProductByMinMaxPriceUseCase"
)
