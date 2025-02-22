package http_errors

import (
	"coworking/internal/spaces/hotdesk"
	meetingroom "coworking/internal/spaces/meeting_room"
	"coworking/internal/spaces/office"

	"github.com/gofiber/fiber/v2"
)

var HTTPErrorMapping = map[error]int{
	hotdesk.ErrInvalidHotDeskNumber:           fiber.StatusBadRequest, // 400
	hotdesk.ErrHotDeskAlreadyExists:           fiber.StatusConflict,   // 409
	meetingroom.ErrInvalidMeetingRoomCapacity: fiber.StatusBadRequest, // 400
	meetingroom.ErrMeetingRoomAlreadyExists:   fiber.StatusConflict,   // 409
	meetingroom.ErrInvalidMeetingRoomName:     fiber.StatusBadRequest, // 400
	office.ErrInvalidOfficeLeasePeriod:        fiber.StatusBadRequest, // 400
	office.ErrInvalidOfficeNumber:             fiber.StatusBadRequest, // 400
	office.ErrOfficeAlreadyExists:             fiber.ErrConflict.Code, // 409
}

func MapDomainErrorToHTTPStatus(err error) int {
	if status, exists := HTTPErrorMapping[err]; exists {
		return status
	}
	return fiber.StatusInternalServerError
}
