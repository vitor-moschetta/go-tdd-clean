package shared

type Mediator struct {
	useCases map[UseCaseKey]UseCase
}

func NewMediator() *Mediator {
	return &Mediator{
		useCases: make(map[UseCaseKey]UseCase),
	}
}

func (m *Mediator) RegisterUseCase(name UseCaseKey, useCase UseCase) {
	m.useCases[name] = useCase
}

func (m *Mediator) GetUseCase(name UseCaseKey) UseCase {
	return m.useCases[name]
}

func (m *Mediator) Execute(name UseCaseKey, input any) Output {
	useCase := m.GetUseCase(name)
	if useCase == nil {
		panic("use case not found")
	}

	return useCase.Execute(input)
}
