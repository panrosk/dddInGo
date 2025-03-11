package models

import (
	"github.com/go-playground/validator/v10"
)

type MembershipDTO struct {
	UserID string `json:"user_id" validate:"omitempty,uuid"`
}

type MembershipValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"rule"`
	Param string `json:"param,omitempty"`
}

var membershipValidator = validator.New()

func (dto *MembershipDTO) Validate() []*MembershipValidationError {
	var errors []*MembershipValidationError

	if err := membershipValidator.Struct(dto); err != nil {
		for _, validationErr := range err.(validator.ValidationErrors) {
			errors = append(errors, &MembershipValidationError{
				Field: validationErr.Field(),
				Tag:   validationErr.Tag(),
				Param: validationErr.Param(),
			})
		}
	}

	return errors
}
