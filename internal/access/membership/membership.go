package membership

import (
	"coworking/internal/core/events"
	"time"

	"github.com/google/uuid"
)

var _ events.EventSourcedEntity[MembershipEventTypes, any] = (*Membership)(nil)

type Membership struct {
	id            string
	userID        string
	active        bool
	createdAt     time.Time
	packages      []Package
	pendingEvents []events.DomainEvent[MembershipEventTypes, any]
}

func NewMembership(id string) *Membership {

	if id == "" {
		id = uuid.New().String()
	}

	return &Membership{
		id:       id,
		packages: []Package{},
	}
}

func (m *Membership) GetID() string {
	return m.id
}

func (m *Membership) UserID() string {
	return m.userID
}

func (m *Membership) When(event events.DomainEvent[MembershipEventTypes, any]) {
	m.pendingEvents = append(m.pendingEvents, event)
	m.Apply(event)
}

func (m *Membership) Apply(event events.DomainEvent[MembershipEventTypes, any]) {
	switch event.Type() {

	case MembershipCreatedEvent:
		version := event.Version()

		if version != 0 {
			return
		}

		payload, ok := event.Payload().(MembershipCreatedPayload)
		if !ok {
			return
		}

		m.userID = payload.UserID
		m.active = true
		m.createdAt = event.OccurredAt()

	case PackageSubscribed:
		payload, ok := event.Payload().(PackageSubscribedPayload)
		if !ok {
			return
		}
		m.packages = append(m.packages, Package{
			packageID: payload.PackageID,
			credits:   payload.Credits,
			startedAt: payload.StartedAt,
			endedAt:   payload.EndedAt,
		})

	default:
	}
}

func CreateMembership(userID string) (*Membership, error) {
	newMembership := NewMembership("")

	event, err := NewMembershipCreatedEvent(newMembership.id, 0, MembershipCreatedPayload{
		UserID: userID,
	})

	if err != nil {
		return nil, err
	}

	newMembership.When(event)

	return newMembership, nil
}

func (m *Membership) SubscribePackage(month int, year int, credits int) error {
	if month < 1 || month > 12 {
		return ErrInvalidMonth
	}

	if credits < 1 {
		return ErrInvalidCredits
	}

	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0).Add(-time.Nanosecond)

	event, err := NewPackageSubscribedEvent(m.id, 0, PackageSubscribedPayload{
		PackageID: uuid.New().String(),
		Credits:   credits,
		StartedAt: startDate,
		EndedAt:   endDate,
	})

	if err != nil {
		return err
	}

	m.When(event)

	return nil
}

func (m *Membership) ReleaseEvents() []events.DomainEvent[MembershipEventTypes, any] {
	released := m.pendingEvents
	m.pendingEvents = nil
	return released
}
