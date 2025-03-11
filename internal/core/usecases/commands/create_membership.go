package commands

import (
	"coworking/internal/access/membership"
	"coworking/internal/core/events"
	"coworking/internal/ports"
)

type CreateMembershipParams struct {
	UserID string
}

type CreateMembershipUseCase struct {
	publisher events.EventPublisher[membership.MembershipEventTypes]
	storage   ports.MembershipRepository
}

func NewCreateMembershipUseCase(publisher events.EventPublisher[membership.MembershipEventTypes], storage ports.MembershipRepository) *CreateMembershipUseCase {
	return &CreateMembershipUseCase{
		publisher: publisher,
		storage:   storage,
	}
}

func (u *CreateMembershipUseCase) Handle(params CreateMembershipParams) error {

	if exists, _ := u.storage.ExistsByUserID(params.UserID); exists {
		return membership.ErrMembershipAlreadyExists
	}

	newMembership, err := membership.CreateMembership(params.UserID)

	if err != nil {
		return err
	}

	newEvents := newMembership.ReleaseEvents()

	for _, e := range newEvents {
		if err := u.publisher.Publish(e); err != nil {
			return err
		}
	}

	return nil

}
