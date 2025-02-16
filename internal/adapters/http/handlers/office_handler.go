package handlers

import (
	"coworking/internal/adapters/http/http_errors"
	"coworking/internal/adapters/http/models"
	"coworking/internal/app/usecases"
	"coworking/internal/app/usecases/commands"
	"coworking/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type OfficeHandler struct {
	commands *usecases.OfficeUsecases
}

func NewOfficeHandler(registerCommand *usecases.OfficeUsecases) *OfficeHandler {
	return &OfficeHandler{commands: registerCommand}
}

func (h *OfficeHandler) RegisterRoutes(app *fiber.App) {
	commandsGroup := app.Group("/offices")
	commandsGroup.Post("/", h.RegisterEntity)
}

func (h *OfficeHandler) RegisterEntity(c *fiber.Ctx) error {
	var req models.OfficeDTO

	if err := c.BodyParser(&req); err != nil {
		return formatErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return formatErrorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	params := commands.RegisterOfficeParams{
		Number:      req.Number,
		LeasePeriod: req.LeasePeriod,
		Status:      req.Status,
	}

	office, err := h.commands.RegisterOffice.Execute(params)
	if err != nil {
		statusCode := http_errors.MapDomainErrorToHTTPStatus(err)
		return formatErrorResponse(c, statusCode, "Failed to register office", err.Error())
	}

	return formatSuccessResponse(c, fiber.StatusOK, "Office registered successfully", office.GetOffice())
}

var _ ports.HttpPort = (*OfficeHandler)(nil)
