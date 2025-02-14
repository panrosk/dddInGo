package ports

import "github.com/gofiber/fiber/v2"

type HttpPort interface {
	RegisterRoutes(app *fiber.App)
	RegisterEntity(ctx *fiber.Ctx) error
}
