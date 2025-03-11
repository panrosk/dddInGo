package events

type EventStrategy[T any] interface {
	CreateEvent(aggregateId string, payload T, version int) DomainEvent[T]
}
