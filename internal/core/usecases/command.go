package usecases

import "coworking/internal/core/usecases/commands"

type Command[T any, R any] interface {
	Handle(params T) (R, error)
}

type HotdeskUsecases struct {
	RegisterHotdesk *commands.RegisterHotdeskUsecase
}

type MeetingRoomUsecases struct {
	RegisterMeetingRoom *commands.RegisterMeetingRoomUsecase
}

type OfficeUsecases struct {
	RegisterOffice *commands.RegisterOfficeUsecase
}

type HotdeskReservationUsecases struct {
	RegisterReservation *commands.ReserveHotdeskUsecase
}

type MeetingRoomReservationUsecases struct {
	RegisterReservation *commands.ReserveMeetingRoomUseCase
}

type CreateMembershipUsecases struct {
	CreateMembership *commands.CreateMembershipUseCase
}
