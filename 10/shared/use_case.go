package shared

type UseCase interface {
	Execute(input interface{}) Output
}
