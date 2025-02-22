package handlers

import (
	"coworking/internal/adapters/http/http_errors"
	"coworking/internal/adapters/http/models"
	"coworking/internal/core/usecases"
	"coworking/internal/core/usecases/commands"
	"coworking/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type HotdeskHandler struct {
	commands *usecases.HotdeskUsecases
}

func NewHotdeskHandler(registerCommand *usecases.HotdeskUsecases) *HotdeskHandler {
	return &HotdeskHandler{commands: registerCommand}
}

func (h *HotdeskHandler) RegisterRoutes(core *fiber.App) {
	commandsGroup := core.Group("/hotdesks")
	commandsGroup.Post("/", h.RegisterEntity)
}

func (h *HotdeskHandler) RegisterEntity(c *fiber.Ctx) error {
	var req models.HotdeskDTO

	if err := c.BodyParser(&req); err != nil {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	params := commands.RegisterHotdeskParams{
		Number: req.Number,
	}

	if err := h.commands.RegisterHotdesk.Handle(params); err != nil {
		statusCode := http_errors.MapDomainErrorToHTTPStatus(err)
		return FormatErrorResponse(c, statusCode, "Failed to register hotdesk", err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

var _ ports.HttpPort = (*HotdeskHandler)(nil)
