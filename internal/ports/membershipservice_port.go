package ports

import (
	"github.com/google/uuid"
	"time"
)

type MembershipResponse struct {
	MembershipID     uuid.UUID
	RemainingCredits int
}

type MembershipService interface {
	CheckMembership(userId uuid.UUID, date time.Time) (*MembershipResponse, error)
}
