package appvalidator

import (
	"github.com/go-playground/validator"
)

type AppValidator struct {
	validator *validator.Validate
}

func NewAppValidator() *AppValidator {
	return &AppValidator{validator: validator.New()}
}

func (a *AppValidator) Validate(i interface{}) error {
	return a.validator.Struct(i)
}
