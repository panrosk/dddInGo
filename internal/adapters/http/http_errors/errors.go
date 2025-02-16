package http_errors

import (
	"coworking/internal/app/domain/domain_errors"
	"github.com/gofiber/fiber/v2"
)

var HTTPErrorMapping = map[error]int{
	domain_errors.ErrInvalidHotDeskNumber:       fiber.StatusBadRequest, // 400
	domain_errors.ErrHotDeskAlreadyExists:       fiber.StatusConflict,   // 409
	domain_errors.ErrInvalidMeetingRoomCapacity: fiber.StatusBadRequest, // 400
	domain_errors.ErrMeetingRoomAlreadyExists:   fiber.StatusConflict,   // 409
	domain_errors.ErrInvalidMeetingRoomName:     fiber.StatusBadRequest, // 400
	domain_errors.ErrInvalidOfficeLeasePeriod:   fiber.StatusBadRequest, // 400
	domain_errors.ErrInvalidOfficeNumber:        fiber.StatusBadRequest, // 400
	domain_errors.ErrOfficeAlreadyExists:        fiber.ErrConflict.Code, // 409
}

func MapDomainErrorToHTTPStatus(err error) int {
	if status, exists := HTTPErrorMapping[err]; exists {
		return status
	}
	return fiber.StatusInternalServerError
}
