package models

import "github.com/go-playground/validator/v10"

type OfficeDTO struct {
	Number      int    `json:"number" validate:"required,gte=0"`                            // Obligatorio, puede ser 0 o mayor
	LeasePeriod int    `json:"leasePeriod,omitempty" validate:"omitempty,gte=0"`            // Opcional, puede ser 0 o mayor si se envía
	Status      string `json:"status,omitempty" validate:"omitempty,oneof=Active Inactive"` // Opcional, "Active" o "Inactive"
}

type OfficeValidationErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"rule"`
	Value       string `json:"value,omitempty"`
}

var officeValidator = validator.New()

func (dto *OfficeDTO) Validate() []*OfficeValidationErrorResponse {
	var errors []*OfficeValidationErrorResponse

	err := officeValidator.Struct(dto)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &OfficeValidationErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}
