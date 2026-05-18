package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateTicketCreate(ticket *TicketCreate) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	ticketToValidate := &TicketCreate{
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

func ValidateTicketEdit(ticket *TicketEdit) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	ticketToValidate := &TicketEdit{
		ID:          ticket.ID,
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
