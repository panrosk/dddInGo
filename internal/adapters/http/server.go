package http

import "github.com/gofiber/fiber/v2"

type Server struct {
	*fiber.App
}

func NewServer() *Server {

	server := &Server{fiber.New(
		fiber.Config{
			ServerHeader: "Coworking",
			AppName:      "Coworking",
		},
	)}

	factory := NewHandlerFactory(server)
	factory.RegisterRoutes()
	return server
}
