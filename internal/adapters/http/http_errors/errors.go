package http_errors

import (
	"coworking/internal/access/membership"
	"coworking/internal/spaces/hotdesk"
	meetingroom "coworking/internal/spaces/meeting_room"
	"coworking/internal/spaces/office"
	"github.com/gofiber/fiber/v2"
)

var HTTPErrorMapping = map[error]int{
	hotdesk.ErrInvalidHotDeskNumber:           fiber.StatusBadRequest,
	hotdesk.ErrHotDeskAlreadyExists:           fiber.StatusConflict,
	meetingroom.ErrInvalidMeetingRoomCapacity: fiber.StatusBadRequest,
	meetingroom.ErrMeetingRoomAlreadyExists:   fiber.StatusConflict,
	meetingroom.ErrInvalidMeetingRoomName:     fiber.StatusBadRequest, // 400
	office.ErrInvalidOfficeLeasePeriod:        fiber.StatusBadRequest, // 400
	office.ErrInvalidOfficeNumber:             fiber.StatusBadRequest, // 400
	office.ErrOfficeAlreadyExists:             fiber.StatusConflict,
	hotdesk.ErrHotDeskAlredyReserved:          fiber.StatusConflict,
	meetingroom.ErrInvalidReservationHour:     fiber.StatusBadRequest,
	meetingroom.ErrInvalidDuration:            fiber.StatusBadRequest,
	hotdesk.ErrHotDeskAlredyReserved:          fiber.StatusConflict,
	membership.ErrMembershipAlreadyExists:     fiber.StatusConflict,
}

func MapDomainErrorToHTTPStatus(err error) int {
	if status, exists := HTTPErrorMapping[err]; exists {
		return status
	}
	return fiber.StatusInternalServerError
}
