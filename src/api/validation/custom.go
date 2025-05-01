package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationError(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, v := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.Property = v.Field()
			el.Tag = v.Tag()
			el.value = v.Param()
			validationErrors = append(validationErrors, el)
		}
		return &validationErrors
	}
	return nil
}
