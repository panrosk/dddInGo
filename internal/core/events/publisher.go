package events

type EventPublisher[T any] interface {
	Publish(e DomainEvent[T, any]) error
}
