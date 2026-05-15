package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateTokenRequest(request *TokenRequest) error {
	//---> Validate the request payload.
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return errors.New(err.Error())
	}

	//----> Return error if any
	return nil
}
