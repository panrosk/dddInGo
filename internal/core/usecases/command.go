package usecases

import "coworking/internal/app/usecases/commands"

type Command[T any, R any] interface {
	Execute(params T) (R, error)
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

type ReservationUsecases struct {
	RegisterReservation *commands.RegisterReservationUsecase
}
