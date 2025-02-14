package models

import "github.com/go-playground/validator/v10"

type HotdeskDTO struct {
	Number int `json:"number" validate:"required,gte=0"`
}

type ValidationErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"rule"`
	Value       string `json:"value,omitempty"`
}

var validate = validator.New()

func (dto *HotdeskDTO) Validate() []*ValidationErrorResponse {
	var errors []*ValidationErrorResponse
	if dto.Number == 0 {
		return errors
	}

	err := validate.Struct(dto)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ValidationErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}
