package commands

type ReserveMeetingRoomParams struct {
	UserId uuid.UUID
	Date   time.Time
}

type ReserveHotdeskUsecase struct {
	storage           ports.HotDeskReservationRepositoryPort
	membershipService ports.MembershipService
}
