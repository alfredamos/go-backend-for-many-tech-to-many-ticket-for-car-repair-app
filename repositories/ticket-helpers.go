package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
)

func toTicketResponse(ticket models.Ticket) TicketResponse {
	return TicketResponse{
		ID:            ticket.ID,
		Title:         ticket.Title,
		Description:   ticket.Description,
		CustomerID:    ticket.CustomerID,
		CustomerName:  ticket.Customer.User.Name,
		CustomerEmail: ticket.Customer.User.Email,
		CustomerPhone: ticket.Customer.User.Phone,
		CustomerImage: ticket.Customer.User.Image,
		CreatedAt:     ticket.CreatedAt,
		UpdatedAt:     ticket.UpdatedAt,
	}
}

func ToTicketResponseList(tickets []models.Ticket) []TicketResponse {
	var ticketResponses []TicketResponse

	for _, ticket := range tickets {
		ticketResponses = append(ticketResponses, toTicketResponse(ticket))
	}
	return ticketResponses
}

func getOneTicketById(id string, t *TicketRepositoryImpl) (*models.Ticket, error) {
	//----> Initialize ticket.
	ticket := new(models.Ticket)

	//----> Fetch ticket by id.
	if err := t.DB.Preload("Customer.User").Where("id = ?", id).First(&ticket).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return ticket, nil
}
