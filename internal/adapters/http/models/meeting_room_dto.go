package models

type MeetingRoomDTO struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

type MeetingRoomValidationError struct {
	FailedField string `json:"field"`
	Tag         string `json:"rule"`
	Value       string `json:"value,omitempty"`
}

func NewMeetingRoomDTO(name string, capacity int) *MeetingRoomDTO {
	return &MeetingRoomDTO{
		Name:     name,
		Capacity: capacity,
	}
}

func (dto *MeetingRoomDTO) Validate() []*MeetingRoomValidationError {
	var errors []*MeetingRoomValidationError

	if dto.Capacity < 0 {
		errors = append(errors, &MeetingRoomValidationError{
			FailedField: "capacity",
			Tag:         "gte",
			Value:       "0",
		})
	}

	return errors
}
