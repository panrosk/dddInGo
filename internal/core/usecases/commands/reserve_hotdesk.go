package commands

import (
	"coworking/internal/common"
	"coworking/internal/ports"
	"coworking/internal/spaces/hotdesk"
	"errors"
	"time"

	"github.com/google/uuid"
)

type ReserveHotdeskParams struct {
	UserId uuid.UUID
	Date   time.Time
}

type ReserveHotdeskUsecase struct {
	storage           ports.RepositoryPort[*hotdesk.HotDeskReservation]
	membershipService ports.MembershipService
}

func NewReserveHotdeskUsecase(storage ports.RepositoryPort[*hotdesk.HotDeskReservation], membershipService ports.MembershipService) *ReserveHotdeskUsecase {
	return &ReserveHotdeskUsecase{
		storage:           storage,
		membershipService: membershipService,
	}
}

func (u *ReserveHotdeskUsecase) Handle(params ReserveHotdeskParams) error {
	// Check if there's already an existing reservation for the user on the specified date
	existingReservations, err := u.storage.FindByFilter(func(r *hotdesk.HotDeskReservation) bool {
		return r.UserId == params.UserId && r.Date.Equal(params.Date)
	})
	if err != nil {
		return err
	}

	if len(existingReservations) > 0 {
		return errors.New("a reservation already exists for this user on the specified date")
	}

	membershipResponse, err := u.membershipService.CheckMembership(params.UserId, params.Date)
	if err != nil {
		return err
	}

	if membershipResponse == nil {
		return errors.New("no membership information found")
	}

	if membershipResponse.RemainingCredits <= 0 {
		return errors.New("reservation cannot be made: no remaining credits in membership")
	}

	newReservation, err := hotdesk.NewHotDeskReservation(params.UserId, params.Date, true)

	if err != nil {
		return err
	}
	err = u.storage.Save(newReservation)
	if err != nil {
		return err
	}

	return nil
}

