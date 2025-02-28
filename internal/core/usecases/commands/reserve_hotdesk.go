package commands

import (
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
	storage           ports.HotDeskReservationRepositoryPort
	membershipService ports.MembershipService
}

func NewReserveHotdeskUsecase(storage ports.HotDeskReservationRepositoryPort, membershipService ports.MembershipService) *ReserveHotdeskUsecase {
	return &ReserveHotdeskUsecase{
		storage:           storage,
		membershipService: membershipService,
	}
}

func (u *ReserveHotdeskUsecase) Handle(params ReserveHotdeskParams) error {
	newReservation, err := createReservation(params.UserId, params.Date, true)
	if err != nil {
		return err
	}

	if u.reservationAlreadyExists(params.UserId, params.Date) {
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

	return u.storage.Save(newReservation)
}

func createReservation(userId uuid.UUID, date time.Time, includedInMembership bool) (*hotdesk.Reservation, error) {
	return hotdesk.NewReservation(userId, date, includedInMembership)
}

func (u *ReserveHotdeskUsecase) reservationAlreadyExists(userId uuid.UUID, reservationDate time.Time) bool {
	existingReservations, err := u.storage.FindByUserIDAndDate(userId, reservationDate)
	return err == nil && len(existingReservations) > 0
}
