package storage

import (
	"coworking/internal/access/membership"
	"coworking/internal/core/events"
	"errors"
)

type InMemoryMembershipRepository struct {
	memberships map[string]*membership.Membership
	events      map[string][]events.DomainEvent[membership.MembershipEvent, any]
}

func NewInMemoryMembershipRepository() *InMemoryMembershipRepository {
	return &InMemoryMembershipRepository{
		memberships: make(map[string]*membership.Membership),
		events:      make(map[string][]events.DomainEvent[membership.MembershipEvent, any]),
	}
}

func (r *InMemoryMembershipRepository) Save(m *membership.Membership) error {
	r.memberships[m.GetID()] = m
	return nil
}

func (r *InMemoryMembershipRepository) SaveEvent(event events.DomainEvent[membership.MembershipEvent, any]) error {
	payload, ok := event.Payload().(membership.MembershipEvent)
	if !ok {
		return errors.New("invalid membership event payload")
	}
	membershipID := payload.ID()
	r.events[membershipID] = append(r.events[membershipID], event)
	return nil
}

func (r *InMemoryMembershipRepository) ExistsByUserID(userID string) (bool, error) {
	for _, m := range r.memberships {
		if m.UserID() == userID {
			return true, nil
		}
	}
	return false, nil
}

func (r *InMemoryMembershipRepository) GetByMembershipID(membershipID string) (*membership.Membership, error) {
	m, ok := r.memberships[membershipID]
	if !ok {
		return nil, errors.New("membership not found")
	}
	return m, nil
}

func (r *InMemoryMembershipRepository) LoadEventsByMembershipID(membershipID string) ([]events.DomainEvent[membership.MembershipEvent, any], error) {
	evts, ok := r.events[membershipID]
	if !ok {
		return nil, errors.New("no events found for membership id")
	}
	return evts, nil
}
