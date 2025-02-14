package ports

import "github.com/gofiber/fiber/v2"

type HttpHotdeskPort interface {
	registerHotdesk(ctx *fiber.Ctx) error
}
