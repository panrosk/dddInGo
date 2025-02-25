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
	core.Group("/meeting-rooms").Post("/", h.RegisterMeetingRoom)
}

func (h *MeetingRoomHandler) RegisterMeetingRoom(c *fiber.Ctx) error {
	req, err := parseAndValidateRequest(c)
	if err != nil {
		return err
	}

	params := commands.RegisterMeetingRoomParams{
		Name:     req.Name,
		Capacity: req.Capacity,
	}

	if err := h.commands.RegisterMeetingRoom.Handle(params); err != nil {
		return handleDomainError(c, err)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func parseAndValidateRequest(c *fiber.Ctx) (*models.MeetingRoomDTO, error) {
	var req models.MeetingRoomDTO
	if err := c.BodyParser(&req); err != nil {
		return nil, FormatErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		return nil, FormatErrorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	return &req, nil
}

func handleDomainError(c *fiber.Ctx, err error) error {
	statusCode := http_errors.MapDomainErrorToHTTPStatus(err)
	return FormatErrorResponse(c, statusCode, "Failed to register meeting room", err.Error())
}

var _ ports.HttpPort = (*MeetingRoomHandler)(nil)
