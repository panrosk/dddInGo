package http

import (
	"coworking/internal/adapters/http/handlers"
	"coworking/internal/adapters/storage"
	"coworking/internal/core/usecases"
	"coworking/internal/core/usecases/commands"
)

type HandlerFactory struct {
	server *Server
}

func NewHandlerFactory(server *Server) *HandlerFactory {
	return &HandlerFactory{server: server}
}

func (f *HandlerFactory) RegisterRoutes() {
	f.registerHotdeskRoutes()
	f.registerMeetingRoomRoutes()
	f.registerOfficeRoutes()
}

func (f *HandlerFactory) registerHotdeskRoutes() {
	repo := storage.NewHotDeskRepository()
	usecase := commands.NewRegisterHotdeskUsecase(repo)
	usecases := usecases.HotdeskUsecases{RegisterHotdesk: usecase}
	handler := handlers.NewHotdeskHandler(&usecases)
	handler.RegisterRoutes(f.server.App)
}

func (f *HandlerFactory) registerMeetingRoomRoutes() {
	repo := storage.NewMeetingRoomRepository()
	usecase := commands.NewRegisterMeetingRoomUsecase(repo)
	usecases := usecases.MeetingRoomUsecases{RegisterMeetingRoom: usecase}
	handler := handlers.NewMeetingRoomHandler(&usecases)
	handler.RegisterRoutes(f.server.App)
}

func (f *HandlerFactory) registerOfficeRoutes() {
	repo := storage.NewOfficeRepository()
	usecase := commands.NewRegisterOfficeUsecase(repo)
	usecases := usecases.OfficeUsecases{RegisterOffice: usecase}
	handler := handlers.NewOfficeHandler(&usecases)
	handler.RegisterRoutes(f.server.App)
}

/* Todo porque se necesita el servicio de membership. */
/**/

/* func (f *HandlerFactory) registerReservationRoutes() { */
/* 	repo := storage.NewHotDeskReservationRepository() */
/* 	usecase := commands.NewReserveHotdeskUsecase(repo) */
/* 	usecases := usecases.ReservationUsecases{RegisterReservation: usecase} */
/* 	handler := handlers.NewReservationHandler(&usecases) */
/* 	handler.RegisterRoutes(f.server.App) */
/* } */
