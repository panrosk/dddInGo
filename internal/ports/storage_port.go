package ports

type RepositoryPort[T any] interface {
	Save(entity T) error
	FindById(id any) (T, error)
	FindAll() ([]T, error)
	FindByFilter(filter func(T) bool) ([]T, error)
}
