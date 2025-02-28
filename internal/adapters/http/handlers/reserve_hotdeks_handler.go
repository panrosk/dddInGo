package handlers

import (
	"coworking/internal/adapters/http/http_errors"
	"coworking/internal/adapters/http/models"
	"coworking/internal/core/usecases"
	"coworking/internal/core/usecases/commands"
	"coworking/internal/ports"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HotdeskReservationHandler struct {
	commands *usecases.HotdeskReservationUsecases
}

func NewReservationHandler(registerCommand *usecases.HotdeskReservationUsecases) *HotdeskReservationHandler {
	return &HotdeskReservationHandler{commands: registerCommand}
}

func (h *HotdeskReservationHandler) RegisterRoutes(core *fiber.App) {
	commandsGroup := core.Group("/reservations")
	commandsGroup.Post("/", h.RegisterEntity)
}

func (h *HotdeskReservationHandler) RegisterEntity(c *fiber.Ctx) error {
	var req models.ReservationDTO

	if err := c.BodyParser(&req); err != nil {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID format", err.Error())
	}

	reservationDate, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Invalid date format", err.Error())
	}

	params := commands.ReserveHotdeskParams{
		UserId: userId,
		Date:   reservationDate,
	}

	if err := h.commands.RegisterReservation.Handle(params); err != nil {
		statusCode := http_errors.MapDomainErrorToHTTPStatus(err)
		return FormatErrorResponse(c, statusCode, "Failed to register reservation", err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

var _ ports.HttpPort = (*HotdeskReservationHandler)(nil)
