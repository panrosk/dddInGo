package ports

type RepositoryPort[T any] interface {
	Save(entity T) error
	FindAll() ([]T, error)
}
