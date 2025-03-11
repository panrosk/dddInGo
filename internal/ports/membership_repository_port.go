package ports

import (
	"coworking/internal/access/membership"
	"coworking/internal/core/events"
)

type MembershipRepository interface {
	Save(membership *membership.Membership) error
	ExistsByUserID(userID string) (bool, error)
	GetByMembershipID(membershipID string) (*membership.Membership, error)
	SaveEvent(event interface{}) error
	LoadEventsByMembershipId(membershipID string) ([]events.DomainEvent[membership.MembershipEventTypes, any], error)
	LoadEventsByUserID(userID string) ([]events.DomainEvent[membership.MembershipEventTypes, any], error)
}
