package commands

import (
	"coworking/internal/access/membership"
	"coworking/internal/core/events"
	"coworking/internal/ports"
)

type SubscribePackageParams struct {
	MembershipID string
	PackageID    string
	Credits      int
	Month        int
	Year         int
}

type SubscribePackageUseCase struct {
	publisher events.EventPublisher[membership.MembershipEventTypes]
	storage   ports.MembershipRepository
}

func NewSubscribePackageUseCase(publisher events.EventPublisher[membership.MembershipEventTypes], storage ports.MembershipRepository) *SubscribePackageUseCase {
	return &SubscribePackageUseCase{
		publisher: publisher,
		storage:   storage,
	}
}

func (u *SubscribePackageUseCase) Handle(params SubscribePackageParams) error {

	currentMembership, err := u.storage.GetByMembershipID(params.MembershipID)

	if err != nil {
		return membership.ErrMembershipNotFound
	}

	err = currentMembership.SubscribePackage(params.Month, params.Month, params.Credits)

	if err != nil {
		return err
	}

	u.storage.Save(currentMembership)

	newEvents := currentMembership.ReleaseEvents()

	for _, e := range newEvents {
		if err := u.publisher.Publish(e); err != nil {
			return err
		}
	}

	return nil
}
