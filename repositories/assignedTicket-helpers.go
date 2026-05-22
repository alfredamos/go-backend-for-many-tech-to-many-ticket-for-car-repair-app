package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
)

func toAssignedTicketResponse(ticket models.AssignedTicket) AssignedTicketResponse {
	//----> Check for existence of ticket and technician.
	if ticket.Ticket.Customer == nil || ticket.Technician.User == nil {
		return AssignedTicketResponse{}
	}

	//----> Send back response.
	return AssignedTicketResponse{
		TechID:          ticket.TechnicianID,
		TicketID:        ticket.TicketID,
		Status:          ticket.Status,
		Completed:       ticket.Completed,
		AssignBy:        ticket.AssignBy,
		AssignAt:        ticket.AssignAt,
		TicketTitle:     ticket.Ticket.Title,
		TicketNotes:     ticket.Ticket.Description,
		CustomerName:    ticket.Ticket.Customer.User.Name,
		CustomerEmail:   ticket.Ticket.Customer.User.Email,
		CustomerAddress: ticket.Ticket.Customer.Address,
		CustomerPhone:   ticket.Ticket.Customer.User.Phone,
		CustomerImage:   ticket.Ticket.Customer.User.Image,
		TechName:        ticket.Technician.User.Name,
		TechEmail:       ticket.Technician.User.Email,
		TechPhone:       ticket.Technician.User.Phone,
		TechSpecialty:   ticket.Technician.Specialty,
		TechImage:       ticket.Technician.User.Image,
	}
}

func toAssignedTicketResponseList(tickets []models.AssignedTicket) []AssignedTicketResponse {
	//----> Check for empty tickets.
	if len(tickets) == 0 {
		return []AssignedTicketResponse{}
	}

	var assignedTicketResponses []AssignedTicketResponse
	for _, ticket := range tickets {

		//----> Check for existence of ticket and technician.
		if ticket.Ticket.Customer != nil && ticket.Technician.User != nil {
			assignedTicketResponses = append(assignedTicketResponses, toAssignedTicketResponse(ticket))
		}

	}
	return assignedTicketResponses
}

func getOneAssignedTicketById(techId, ticketId string, t *AssignedTicketRepositoryImpl) (models.AssignedTicket, error) {
	//----> Initialize ticket.
	ticket := new(models.AssignedTicket)

	//----> Fetch ticket by id.
	if err := t.DB.Preload("Ticket.Customer.User").Preload("Technician.User").Where("technician_id = ? AND ticket_id = ?", techId, ticketId).First(&ticket).Error; err != nil {
		return models.AssignedTicket{}, errors.New(err.Error())
	}

	//----> Send back response.
	return *ticket, nil
}
