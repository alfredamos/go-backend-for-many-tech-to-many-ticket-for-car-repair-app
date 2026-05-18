package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateCustomerCreate(customer *CustomerCreate) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	customerToValidate := &CustomerCreate{
		Address: customer.Address,
		Notes:   customer.Notes,
		Active:  customer.Active,
		Status:  customer.Status,
		UserID:  customer.UserID,
	}

	//----> Check for validation error
	if err := validate.Struct(customerToValidate); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil
}

func ValidateCustomerEdit(customer *CustomerEdit) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	customerToValidate := &CustomerEdit{
		ID:      customer.ID,
		Address: customer.Address,
		Notes:   customer.Notes,
		Active:  customer.Active,
		Status:  customer.Status,
		UserID:  customer.UserID,
	}

	//----> Check for validation error
	if err := validate.Struct(customerToValidate); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil
}
