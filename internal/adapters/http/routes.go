package http

import (
	"coworking/internal/adapters/http/handlers"
	"coworking/internal/adapters/storage"
	"coworking/internal/app/usecases"
	"coworking/internal/app/usecases/commands"
)

func (s *Server) RegisterRoutes() {
	hotdeskRepo := storage.NewHotDeskRepository()
	hotdeskRegisterCommand := commands.NewRegisterHotdeskUsecase(hotdeskRepo)
	hotdeskUsecases := usecases.HotdeskUsecases{RegisterHotdesk: hotdeskRegisterCommand}
	hotdeskHandler := handlers.NewHotdeskHandler(&hotdeskUsecases)
	hotdeskHandler.RegisterRoutes(s.App)

	meetingRoomRepo := storage.NewMeetingRoomRepository()
	meetingRoomRegisterCommand := commands.NewRegisterMeetingRoomUsecase(meetingRoomRepo)
	meetingRoomUsecases := usecases.MeetingRoomUsecases{RegisterMeetingRoom: meetingRoomRegisterCommand}
	meetingRoomHandler := handlers.NewMeetingRoomHandler(&meetingRoomUsecases)
	meetingRoomHandler.RegisterRoutes(s.App)

	officeRepo := storage.NewOfficeRepository()
	officeRegisterCommand := commands.NewRegisterOfficeUsecase(officeRepo)
	officeUsecases := usecases.OfficeUsecases{RegisterOffice: officeRegisterCommand}
	officeHandler := handlers.NewOfficeHandler(&officeUsecases)
	officeHandler.RegisterRoutes(s.App)

}
