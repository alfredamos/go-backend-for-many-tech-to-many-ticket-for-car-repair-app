package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateTechnicianCreate(tech *TechnicianCreate) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	technicianToValidate := &TechnicianCreate{
		Specialty: tech.Specialty,
		UserID:    tech.UserID,
	}

	//----> Check for validation error
	if err := validate.Struct(technicianToValidate); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil
}

func ValidateTechnicianEdit(tech *TechnicianEdit) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	technicianToValidate := &TechnicianEdit{
		ID:        tech.ID,
		Specialty: tech.Specialty,
		UserID:    tech.UserID,
	}

	//----> Check for validation error
	if err := validate.Struct(technicianToValidate); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil
}
