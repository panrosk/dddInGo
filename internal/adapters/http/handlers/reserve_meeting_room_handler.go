package handlers

import (
	"coworking/internal/adapters/http/models"
	"coworking/internal/core/usecases"
	"coworking/internal/core/usecases/commands"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ReserveMeetingRoomHandler struct {
	reserveMeetingRoomUsecase *usecases.MeetingRoomReservationUsecases
	reserveHotdeskUsecase     *usecases.HotdeskReservationUsecases
}

func NewReserveMeetingRoomHandler(reserveMeetingRoomUsecase *usecases.MeetingRoomReservationUsecases, reserveHotdeskUsecase *usecases.HotdeskReservationUsecases) *ReserveMeetingRoomHandler {
	return &ReserveMeetingRoomHandler{
		reserveMeetingRoomUsecase: reserveMeetingRoomUsecase,
		reserveHotdeskUsecase:     reserveHotdeskUsecase,
	}
}

func (h *ReserveMeetingRoomHandler) RegisterRoutes(core *fiber.App) {
	commandsGroup := core.Group("/meeting-room-reservations")
	commandsGroup.Post("/", h.RegisterEntity)
}

func (h *ReserveMeetingRoomHandler) RegisterEntity(c *fiber.Ctx) error {
	var req models.MeetingRoomReservation

	if err := c.BodyParser(&req); err != nil {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	params := commands.ReserveMeetingRoomParams{
		MeetingRoom: req.MeetingRoomId,
		Date:        req.Date,
		Hour:        req.Hour,
		Duration:    req.Duration,
		UserId:      req.UserId,
	}

	if err := h.reserveMeetingRoomUsecase.RegisterReservation.Handle(params); err != nil {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Failed to register reservation", err.Error())
	}

	userId, err := uuid.Parse(req.UserId)

	if err != nil {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID format", err.Error())
	}

	reservationDate, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		return FormatErrorResponse(c, fiber.StatusBadRequest, "Invalid date format", err.Error())
	}

	paramsHotdesk := commands.ReserveHotdeskParams{
		UserId: userId,
		Date:   reservationDate,
	}

	var hotdeskRegistered bool = true

	err = h.reserveHotdeskUsecase.RegisterReservation.Handle(paramsHotdesk)

	if err != nil {
		hotdeskRegistered = false
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"hotdesk_registered": hotdeskRegistered,
	})

}
