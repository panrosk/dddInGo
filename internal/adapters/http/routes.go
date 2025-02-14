package http

import (
	"coworking/internal/adapters/http/handlers"
	"coworking/internal/adapters/storage"
	"coworking/internal/app/usecases"
	"coworking/internal/app/usecases/commands"
)

func (s *Server) RegisterRoutes() {
	// Hotdesk setup
	hotdeskRepo := storage.NewHotDeskRepository()
	hotdeskRegisterCommand := commands.NewRegisterHotdeskUsecase(hotdeskRepo)
	hotdeskUsecases := usecases.HotdeskUsecases{RegisterHotdesk: hotdeskRegisterCommand}
	hotdeskHandler := handlers.NewHotdeskHandler(&hotdeskUsecases)
	hotdeskHandler.RegisterRoutes(s.App)

	// Meeting Room setup
	meetingRoomRepo := storage.NewMeetingRoomRepository()
	meetingRoomRegisterCommand := commands.NewRegisterMeetingRoomUsecase(meetingRoomRepo)
	meetingRoomUsecases := usecases.MeetingRoomUsecases{RegisterMeetingRoom: meetingRoomRegisterCommand}
	meetingRoomHandler := handlers.NewMeetingRoomHandler(&meetingRoomUsecases)
	meetingRoomHandler.RegisterRoutes(s.App)
}
