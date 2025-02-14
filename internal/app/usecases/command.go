package usecases

type Command[T any, R any] interface {
	Execute(params T) (R, error)
}
