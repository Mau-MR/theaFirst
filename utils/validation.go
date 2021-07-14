package utils

import (
	"fmt"
	"github.com/go-playground/validator"
	"net/http"
)

//ALL THE CREDITS TO NICHOLAS JACKSON <3

// Validation contains
type Validation struct {
	val *validator.Validate
}

// NewValidation creates a new Validation type
func NewValidation() *Validation {
	validate := validator.New()
	//TODO: ADD HERE CUSTOM VALIDATIONS
	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	var returnErrs []ValidationError
	if errs, ok := v.val.Struct(i).(validator.ValidationErrors); ok {
		if errs != nil {
			for _, err := range errs {
				if fe, ok := err.(validator.FieldError); ok {
					ve := ValidationError{fe}
					returnErrs = append(returnErrs, ve)
				}
			}
		}
	}
	return returnErrs
}
func (v *Validation) ValidateRequest(i interface{}, rw http.ResponseWriter) error {
	errs := v.Validate(i)
	if len(errs) != 0 {
		// return the validation messages as an array
		rw.WriteHeader(http.StatusUnprocessableEntity)
		ToJSON(&ValErrors{Messages: errs.Errors()}, rw)
		return fmt.Errorf("Error with the  fields provided")
	}
	return nil
}
