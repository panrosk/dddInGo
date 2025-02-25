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
	commandsGroup.Post("/", h.RegisterHotdesk)
}

func (h *HotdeskHandler) RegisterHotdesk(c *fiber.Ctx) error {
	req, err := parseRequest(c)
	if err != nil {
		return err
	}

	if err := h.processRegisterHotdesk(c, req); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func parseRequest(c *fiber.Ctx) (models.HotdeskDTO, error) {
	var req models.HotdeskDTO

	if err := c.BodyParser(&req); err != nil {
		return req, FormatErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		return req, FormatErrorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	return req, nil
}

func (h *HotdeskHandler) processRegisterHotdesk(c *fiber.Ctx, req models.HotdeskDTO) error {
	params := commands.RegisterHotdeskParams{
		Number: req.Number,
	}

	if err := h.commands.RegisterHotdesk.Handle(params); err != nil {
		statusCode := http_errors.MapDomainErrorToHTTPStatus(err)
		return FormatErrorResponse(c, statusCode, "Failed to register hotdesk", err.Error())
	}

	return nil
}

var _ ports.HttpPort = (*HotdeskHandler)(nil)
