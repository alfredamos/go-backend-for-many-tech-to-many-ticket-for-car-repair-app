package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func validateTicket(ticket *Ticket) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	ticketToValidate := &Ticket{
		Title:       ticket.Title,
		Description: ticket.Description,
		CustomerID:  ticket.CustomerID,
	}

	//----> Check for validation error
	if err := validate.Struct(ticketToValidate); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil
}
