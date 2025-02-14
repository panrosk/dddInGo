package http

import (
	"coworking/internal/adapters/http/handlers"
	"coworking/internal/adapters/storage"
	"coworking/internal/app/usecases"
	"coworking/internal/app/usecases/commands"
)

func (s *Server) RegisterRoutes() {
	repo := storage.NewHotDeskRepository()
	registerCommand := commands.NewRegisterHotdeskUsecase(repo)
	usecases_commands := usecases.HotdeskUsecases{RegisterHotdesk: registerCommand}
	hotdeskHandler := handlers.NewHotdeskHandler(&usecases_commands)
	hotdeskHandler.RegisterRoutes(s.App)

}
