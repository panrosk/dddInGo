package handlers

import (
	"coworking/internal/adapters/http/http_errors"
	"coworking/internal/adapters/http/models"
	"coworking/internal/app/usecases"
	"coworking/internal/app/usecases/commands"
	"coworking/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type HotdeskHandler struct {
	commands *usecases.HotdeskUsecases
}

func NewHotdeskHandler(registerCommand *usecases.HotdeskUsecases) *HotdeskHandler {
	return &HotdeskHandler{commands: registerCommand}
}

func (h *HotdeskHandler) RegisterRoutes(app *fiber.App) {
	commandsGroup := app.Group("/hotdesks")
	commandsGroup.Post("/", h.RegisterEntity)
}

func (h *HotdeskHandler) RegisterEntity(c *fiber.Ctx) error {
	var req models.HotdeskDTO

	if err := c.BodyParser(&req); err != nil {
		return formatErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return formatErrorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	params := commands.RegisterHotdeskParams{
		Number: req.Number,
	}

	hotdesk, err := h.commands.RegisterHotdesk.Execute(params)
	if err != nil {
		statusCode := http_errors.MapDomainErrorToHTTPStatus(err)
		return formatErrorResponse(c, statusCode, "Failed to register hotdesk", err.Error())
	}

	return formatSuccessResponse(c, fiber.StatusOK, "Hotdesk registered successfully", hotdesk.GetHotdesk())
}

var _ ports.HttpPort = (*HotdeskHandler)(nil)

var _ ports.HttpPort = (*HotdeskHandler)(nil)
