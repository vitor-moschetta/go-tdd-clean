package shared

const (
	CreateCategoryUseCase  = "CreateCategoryUseCase"
	GetCategoryByIDUseCase = "GetCategoryByIDUseCase"
)

const (
	CreateProductUseCase           = "CreateProductUseCase"
	GetProductByMinMaxPriceUseCase = "GetProductByMinMaxPriceUseCase"
)

type Mediator struct {
	useCases map[string]UseCase
}

func NewMediator() *Mediator {
	return &Mediator{
		useCases: make(map[string]UseCase),
	}
}

func (m *Mediator) RegisterUseCase(name string, useCase UseCase) {
	m.useCases[name] = useCase
}

func (m *Mediator) GetUseCase(name string) UseCase {
	return m.useCases[name]
}

func (m *Mediator) Execute(name string, input interface{}) Output {
	useCase := m.GetUseCase(name)
	if useCase == nil {
		panic("use case not found")
	}

	return useCase.Execute(input)
}
