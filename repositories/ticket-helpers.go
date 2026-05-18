package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
)

func toTicketResponse(ticket models.Ticket, user *models.User) TicketResponse {
	return TicketResponse{
		ID:            ticket.ID,
		Title:         ticket.Title,
		Description:   ticket.Description,
		CustomerID:    ticket.CustomerID,
		CustomerName:  user.Name,
		CustomerEmail: user.Email,
		CustomerPhone: user.Phone,
		CustomerImage: user.Image,
		CreatedAt:     ticket.CreatedAt,
		UpdatedAt:     ticket.UpdatedAt,
	}
}

func ToTicketResponseList(tickets []models.Ticket, t *TicketRepositoryImpl) []TicketResponse {
	var ticketResponses []TicketResponse

	for _, ticket := range tickets {
		user, err := getOneUserByIdInTicket(ticket.CustomerID, t)

		//----> Check for error.
		if err != nil {
			continue
		}

		ticketResponses = append(ticketResponses, toTicketResponse(ticket, user))
	}
	return ticketResponses
}

func getOneTicketById(id string, t *TicketRepositoryImpl) (*models.Ticket, error) {
	//----> Initialize ticket.
	ticket := new(models.Ticket)

	//----> Fetch ticket by id.
	if err := t.DB.Where("id = ?", id).First(&ticket).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return ticket, nil
}

func getOneUserByIdInTicket(customerId string, c *TicketRepositoryImpl) (*models.User, error) {
	//----> Initialize user.
	customer := new(models.Customer)
	user := new(models.User)

	//----> Fetch customer by id.
	if err := c.DB.Where("id = ?", customerId).First(&customer).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Fetch user by id.
	if err := c.DB.Where("id = ?", customer.UserID).First(&user).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return user, nil
}
