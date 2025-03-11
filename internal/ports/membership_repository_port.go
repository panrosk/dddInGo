package ports

import "coworking/internal/access/membership"

type MembershipRepository interface {
	Save(membership *membership.Membership) error
	ExistsByUserID(userID string) (bool, error)
	GetByMembershipID(membershipID string) (*membership.Membership, error)
}
