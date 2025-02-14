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

	server.RegisterRoutes()
	return server
}
