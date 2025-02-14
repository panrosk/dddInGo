package http

import "github.com/gofiber/fiber/v2"

func (s *Server) RegisterRoutes() {
	s.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
