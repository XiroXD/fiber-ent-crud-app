package validator

import (
	"github.com/go-playground/validator/v10"
)

var v = validator.New()

func GetValidator() *validator.Validate {
	return v
}
