package events

type EventSourcedEntity[T, P any] interface {
	Apply(event DomainEvent[T, P])
	When(event DomainEvent[T, P])
	ReleaseEvents() []DomainEvent[T, P]
	GetID() string
}
