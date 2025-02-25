package usecases

type Query[T any, R any] interface {
	Handle(params T) (R, error)
}
