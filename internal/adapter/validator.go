package adapter

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type GvValidator struct {
	gvValidate *validator.Validate
}

func NewValidator() *GvValidator {
	return &GvValidator{gvValidate: validator.New()}
}

func (v *GvValidator) Validate(i any) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := v.gvValidate.Struct(i)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors 
}

type defaultValidator struct {
	adapter *Adapter
}

func DefaultValidator() Option {
	return &defaultValidator{}
}

func (v *defaultValidator) Start(a *Adapter) {
	a.Validator = NewValidator()
	v.adapter = a
}

func (v *defaultValidator) Close() error {
	return nil
}
