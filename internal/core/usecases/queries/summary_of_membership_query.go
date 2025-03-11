package queries

import (
	"coworking/internal/access/membership"
	"coworking/internal/ports"
)

type SummaryOfMembershipQueryParams struct {
	UserID string
}

type SummaryOfMembershipUseCase struct {
	MembershipRepository ports.MembershipRepository
}

type SummaryOfMembershipQueryResult struct {
	UserID       string
	MembershipID string
	TotalCredits int
}

func (u *SummaryOfMembershipUseCase) Handle(params SummaryOfMembershipQueryParams) (SummaryOfMembershipQueryResult, error) {
	events, err := u.MembershipRepository.LoadEventsByUserID(params.UserID)

	if err != nil {
		return SummaryOfMembershipQueryResult{}, err
	}

	newMembership, err := membership.CreateMembership(params.UserID)

	if err != nil {
		return SummaryOfMembershipQueryResult{}, err
	}

	for _, e := range events {
		newMembership.When(e)
	}

	return SummaryOfMembershipQueryResult{
		UserID:       newMembership.UserID(),
		MembershipID: newMembership.GetID(),
		TotalCredits: newMembership.TotalCredits(),
	}, nil

}
