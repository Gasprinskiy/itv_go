package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct[T any](input T) error {
	return validate.Struct(input)
}
