package data

import (
	"fmt"
	"github.com/go-playground/validator"
	"regexp"
)

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key:'%s' Error: Field validation for '%s' failed on the %s tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice in to string slice
func (v ValidationErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// validation contains
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new validation type
func NewValidation() *Validation {
	validate := validator.New()
	_ = validate.RegisterValidation("sku", validateSKU)
	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	var returnErrs []ValidationError

	if errs, ok := v.validate.Struct(i).(validator.ValidationErrors); ok {
		if len(errs) == 0 {
			return nil
		}
		for _, err := range errs {
			// cast the FieldError into our ValidationError and append to the slice
			ve := ValidationError{err.(validator.FieldError)}
			returnErrs = append(returnErrs, ve)
		}
	}

	return returnErrs
}

func validateSKU(fl validator.FieldLevel) bool {
	// SKU must be in format abc-abc-abc
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	sku := re.FindAllString(fl.Field().String(), -1)

	if len(sku) == 1 {
		return true
	}

	return false
}
