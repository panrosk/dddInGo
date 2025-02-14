package handlers

import (
	"coworking/internal/adapters/http/http_errors"
	"coworking/internal/adapters/http/models"
	"coworking/internal/app/usecases"
	"coworking/internal/app/usecases/commands"
	"coworking/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type MeetingRoomHandler struct {
	commands *usecases.MeetingRoomUsecases
}

func NewMeetingRoomHandler(registerCommand *usecases.MeetingRoomUsecases) *MeetingRoomHandler {
	return &MeetingRoomHandler{commands: registerCommand}
}

func (h *MeetingRoomHandler) RegisterRoutes(app *fiber.App) {
	commandsGroup := app.Group("/meeting-rooms")
	commandsGroup.Post("/", h.RegisterEntity)
}

func (h *MeetingRoomHandler) RegisterEntity(c *fiber.Ctx) error {
	var req models.MeetingRoomDTO

	if err := c.BodyParser(&req); err != nil {
		return formatErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return formatErrorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	params := commands.RegisterMeetingRoomParams{
		Name:     req.Name,
		Capacity: req.Capacity,
	}

	meetingRoom, err := h.commands.RegisterMeetingRoom.Execute(params)
	if err != nil {
		statusCode := http_errors.MapDomainErrorToHTTPStatus(err)
		return formatErrorResponse(c, statusCode, "Failed to register meeting room", err.Error())
	}

	return formatSuccessResponse(c, fiber.StatusOK, "Meeting room registered successfully", meetingRoom.GetMeetingRoom())
}

var _ ports.HttpPort = (*MeetingRoomHandler)(nil)
