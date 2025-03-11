package events

import "time"

type Event[T any] interface {
	ID() string
	AggregateID() string
	Type() T
	OccurredAt() time.Time
	Version() int
}

type DomainEvent[T, P any] interface {
	Event[T]
	Payload() P
}
