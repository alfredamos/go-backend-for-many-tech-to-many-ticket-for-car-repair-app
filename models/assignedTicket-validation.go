package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateAssignedTicketCreate(ticket *AssignedTicketCreate) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	assignedTicketToValidate := &AssignedTicketCreate{
		TicketID:     ticket.TicketID,
		TechnicianID: ticket.TechnicianID,
		AssignBy:     ticket.AssignBy,
	}

	//----> Check for validation error
	if err := validate.Struct(assignedTicketToValidate); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil
}

func ValidateAssignedTicketEdit(ticket *AssignedTicketEdit) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	assignedTicketToValidate := &AssignedTicketEdit{
		TicketID:     ticket.TicketID,
		TechnicianID: ticket.TechnicianID,
		AssignBy:     ticket.AssignBy,
		Completed:    ticket.Completed,
		Status:       ticket.Status,
	}

	//----> Check for validation error
	if err := validate.Struct(assignedTicketToValidate); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil
}
