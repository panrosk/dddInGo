package models

import (
	"github.com/go-playground/validator/v10"
)

type ReservationDTO struct {
	UserId string `json:"user_id" validate:"required,uuid"`                            // Obligatorio, debe ser un UUID válido
	Date   string `json:"date" validate:"required,datetime=2006-01-02T15:04:05Z07:00"` // Obligatorio, debe seguir formato ISO 8601
}

type ReservationValidationErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"rule"`
	Value       string `json:"value,omitempty"`
}

var reservationValidator = validator.New()

func (dto *ReservationDTO) Validate() []*ReservationValidationErrorResponse {
	var errors []*ReservationValidationErrorResponse

	err := reservationValidator.Struct(dto)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ReservationValidationErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}
