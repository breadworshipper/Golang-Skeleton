package adapter

import (
	"mm-pddikti-cms/pkg/response"

	"github.com/go-playground/validator/v10"
)

type GvValidator struct {
	gvValidate *validator.Validate
}

func NewValidator() *GvValidator {
	return &GvValidator{gvValidate: validator.New()}
}

func (v *GvValidator) Validate(i any) []response.ErrorValidation {
	validationErrors := []response.ErrorValidation{}

	errs := v.gvValidate.Struct(i)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem response.ErrorValidation

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
