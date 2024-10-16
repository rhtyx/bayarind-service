package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ParseValidationError(err error) map[string]string {
	errFields := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {

		errFields[err.Field()] = fmt.Sprintf(
			"failed '%s' tag check (value '%s' is not valid)", err.Tag(), err.Value(),
		)
	}

	return errFields
}
