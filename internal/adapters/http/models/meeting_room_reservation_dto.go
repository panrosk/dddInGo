package models

import (
	"github.com/go-playground/validator/v10"
)

type MeetingRoomReservation struct {
	MeetingRoomId string `json:"meeting_room_id" validate:"required"` // Obligatorio, identificador de la MeetingRoom a reservar
	Date          string `json:"date" validate:"required"`            // Obligatorio, debe ser una fecha válida
	Hour          int    `json:"hour" validate:"required"`            // Obligatorio, representa la hora de inicio
	Duration      int    `json:"duration" validate:"required"`        // Obligatorio, duración de la reserva en horas
	UserId        string `json:"user_id" validate:"required"`         // Obligatorio, id del usuario que reserva
}

type MeetingRoomReservationValidationErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"rule"`
	Value       string `json:"value,omitempty"`
}

var meetingRoomReservationValidator = validator.New()

func (dto *MeetingRoomReservation) Validate() []*MeetingRoomReservationValidationErrorResponse {
	var errors []*MeetingRoomReservationValidationErrorResponse

	err := meetingRoomReservationValidator.Struct(dto)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &MeetingRoomReservationValidationErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}
