package handlers

import (
	"coworking/internal/adapters/http/http_errors"
	"coworking/internal/adapters/http/models"
	"coworking/internal/core/usecases"
	"coworking/internal/core/usecases/commands"
	"coworking/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type OfficeHandler struct {
	commands *usecases.OfficeUsecases
}

func NewOfficeHandler(registerCommand *usecases.OfficeUsecases) *OfficeHandler {
	return &OfficeHandler{commands: registerCommand}
}

func (h *OfficeHandler) RegisterRoutes(core *fiber.App) {
	commandsGroup := core.Group("/offices")
	commandsGroup.Post("/", h.RegisterEntity)
}

func (h *OfficeHandler) RegisterEntity(c *fiber.Ctx) error {
	var req models.OfficeDTO

	if err := c.BodyParser(&req); err != nil {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	params := commands.RegisterOfficeParams{
		Number:      req.Number,
		LeasePeriod: req.LeasePeriod,
		Status:      req.Status,
	}

	if err := h.commands.RegisterOffice.Handle(params); err != nil {
		statusCode := http_errors.MapDomainErrorToHTTPStatus(err)
		return FormatErrorResponse(c, statusCode, "Failed to register office", err.Error())
	}

	return c.SendStatus(fiber.StatusCreated) // 201 Created
}

var _ ports.HttpPort = (*OfficeHandler)(nil)
