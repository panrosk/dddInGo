package storage

import (
	"coworking/internal/access/membership"
	"coworking/internal/ports"
)

type MembershipRepository struct {
	store          map[string]*membership.Membership
	userIndex      map[string]string
	aggregateIndex map[string]string
}

func NewMembershipRepository() ports.MembershipRepository {
	return &MembershipRepository{
		store:          make(map[string]*membership.Membership),
		userIndex:      make(map[string]string),
		aggregateIndex: make(map[string]string),
	}
}

func (r *MembershipRepository) Save(m *membership.Membership) error {
	r.store[m.GetID()] = m
	r.userIndex[m.UserID()] = m.GetID()
	r.aggregateIndex[m.GetID()] = m.GetID()
	return nil
}

func (r *MembershipRepository) ExistsByUserID(userID string) (bool, error) {
	_, ok := r.userIndex[userID]
	return ok, nil
}

func (r *MembershipRepository) ExistsByMembershipID(membershipID string) (bool, error) {
	_, ok := r.aggregateIndex[membershipID]
	return ok, nil
}
