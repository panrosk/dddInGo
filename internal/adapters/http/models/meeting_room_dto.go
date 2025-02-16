package models

import "github.com/go-playground/validator/v10"

type MeetingRoomDTO struct {
	Name     string `json:"name" validate:"required"`  // Obligatorio
	Capacity int    `json:"capacity" validate:"gte=0"` // Debe ser >= 0
}

type MeetingRoomValidationErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"rule"`
	Value       string `json:"value,omitempty"`
}

var meetingRoomValidator = validator.New()

func NewMeetingRoomDTO(name string, capacity int) *MeetingRoomDTO {
	return &MeetingRoomDTO{
		Name:     name,
		Capacity: capacity,
	}
}

func (dto *MeetingRoomDTO) Validate() []*MeetingRoomValidationErrorResponse {
	var errors []*MeetingRoomValidationErrorResponse

	err := meetingRoomValidator.Struct(dto)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &MeetingRoomValidationErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}

	return errors
}
