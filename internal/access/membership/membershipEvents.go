package membership

import (
	"time"

	"github.com/google/uuid"
)

type MembershipEventTypes string

const (
	MembershipCreatedEvent MembershipEventTypes = "membership_created"
	PackageSubscribed      MembershipEventTypes = "package_subscribed"
)

type MembershipEvent struct {
	id          string
	aggregateID string
	occurredAt  time.Time
	version     int
	eventType   MembershipEventTypes
	payload     any
}

func (me MembershipEvent) ID() string {
	return me.id
}

func (me MembershipEvent) AggregateID() string {
	return me.aggregateID
}

func (me MembershipEvent) Type() MembershipEventTypes {
	return me.eventType
}

func (me MembershipEvent) OccurredAt() time.Time {
	return me.occurredAt
}

func (me MembershipEvent) Version() int {
	return me.version
}

func (me MembershipEvent) Payload() any {
	return me.payload
}

func NewMembershipCreatedEvent(
	aggregateID string,
	version int,
	payload MembershipCreatedPayload,
) (MembershipEvent, error) {
	event := MembershipEvent{
		id:          uuid.New().String(),
		aggregateID: aggregateID,
		occurredAt:  time.Now(),
		version:     version + 1,
		eventType:   MembershipCreatedEvent,
		payload:     payload,
	}
	return event, nil
}

func NewPackageSubscribedEvent(
	aggregateID string,
	version int,
	payload PackageSubscribedPayload,
) (MembershipEvent, error) {
	event := MembershipEvent{
		id:          uuid.New().String(),
		aggregateID: aggregateID,
		occurredAt:  time.Now(),
		version:     version + 1,
		eventType:   PackageSubscribed,
		payload:     payload,
	}
	return event, nil
}
