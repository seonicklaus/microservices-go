package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidationError wraps the validators FieldError so we do not
// expose this to our code
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

// Errors converts the slice into a slice of strings
func (v ValidationErrors) Errors() []string {
	errs := []string{}

	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validation contains
type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError

	for _, err := range errs {
		// cast the FieldError into our validationError and append to the slice
		ve := ValidationError{err}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

// ValidateSKU
func validateSKU(fl validator.FieldLevel) bool {
	// SKU must be in the format abc-defg-hijkl
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := reg.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}
