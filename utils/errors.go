package utils

import (
	"fmt"
	"github.com/go-playground/validator"
)

type GenericError struct {
	Message string `json:"message"`
}
type ValErrors struct {
	Messages []string `json:"messages"`
}

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}
