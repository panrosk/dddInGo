package handlers

import (
	"coworking/internal/adapters/http/http_errors"
	"coworking/internal/adapters/http/models"
	"coworking/internal/core/usecases"
	"coworking/internal/core/usecases/commands"
	"coworking/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type MeetingRoomHandler struct {
	commands *usecases.MeetingRoomUsecases
}

func NewMeetingRoomHandler(registerCommand *usecases.MeetingRoomUsecases) *MeetingRoomHandler {
	return &MeetingRoomHandler{commands: registerCommand}
}

func (h *MeetingRoomHandler) RegisterRoutes(core *fiber.App) {
	commandsGroup := core.Group("/meeting-rooms")
	commandsGroup.Post("/", h.RegisterEntity)
}

func (h *MeetingRoomHandler) RegisterEntity(c *fiber.Ctx) error {
	var req models.MeetingRoomDTO

	if err := c.BodyParser(&req); err != nil {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	params := commands.RegisterMeetingRoomParams{
		Name:     req.Name,
		Capacity: req.Capacity,
	}

	if err := h.commands.RegisterMeetingRoom.Handle(params); err != nil {
		statusCode := http_errors.MapDomainErrorToHTTPStatus(err)
		return FormatErrorResponse(c, statusCode, "Failed to register meeting room", err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

var _ ports.HttpPort = (*MeetingRoomHandler)(nil)
