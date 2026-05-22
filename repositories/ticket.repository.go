package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

	"gorm.io/gorm"
)

type TicketRepositoryInt interface {
	CreateTicket(ticketCreate *models.TicketCreate) (ResponseMessage, error)
	DeleteTicketById(id string) (ResponseMessage, error)
	EditTicketById(id string, ticketEdit *models.TicketEdit) (ResponseMessage, error)
	GetAllTickets() (TicketResponse, error)
	GetTicketById(id string) (TicketResponse, error)
	GetTicketsByCustomerId(customerId string) (TicketResponse, error)
}

type TicketRepositoryImpl struct {
	DB *gorm.DB
}

func NewTicketRepositoryImpl(DB *gorm.DB) *TicketRepositoryImpl {
	return &TicketRepositoryImpl{DB: DB}
}

func (t *TicketRepositoryImpl) CreateTicket(ticketCreate *models.TicketCreate) (ResponseMessage, error) {
	//----> Validate ticket input.
	if err := models.ValidateTicketCreate(ticketCreate); err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Initialize ticket.
	ticket := models.Ticket{
		Title:       ticketCreate.Title,
		Description: ticketCreate.Description,
		CustomerID:  ticketCreate.CustomerID,
	}

	//----> Create ticket.
	if err := t.DB.Create(&ticket).Error; err != nil {
		return ResponseMessage{}, err
	}

	//----> Send back response.
	return NewResponseMessage("Ticket created successfully!", 201, "Success"), nil

}

func (t *TicketRepositoryImpl) DeleteTicketById(id string) (ResponseMessage, error) {
	//----> Check for existence of ticket.
	ticket, err := getOneTicketById(id, t)

	//----> Check for error.
	if err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Delete ticket.
	if err := t.DB.Delete(ticket).Error; err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Send back response.
	return NewResponseMessage("Ticket deleted successfully!", 200, "Success"), nil
}

func (t *TicketRepositoryImpl) EditTicketById(id string, ticketEdit *models.TicketEdit) (ResponseMessage, error) {
	//----> Check for existence of ticket.
	_, err := getOneTicketById(id, t)

	//----> Check for error.
	if err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Validate ticket input.
	if err := models.ValidateTicketEdit(ticketEdit); err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Initialize ticket.
	ticket := models.Ticket{
		ID:          id,
		Title:       ticketEdit.Title,
		Description: ticketEdit.Description,
		CustomerID:  ticketEdit.CustomerID,
	}

	//----> Update ticket.
	if err := t.DB.Updates(&ticket).Error; err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Send back response.
	return NewResponseMessage("Ticket updated successfully!", 200, "Success"), nil
}

func (t *TicketRepositoryImpl) GetTicketById(id string) (TicketResponse, error) {
	//----> Fetch the ticket with the giving id.
	ticket, err := getOneTicketById(id, t)

	//----> Check for error.
	if err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toTicketResponse(*ticket), nil

}

func (t *TicketRepositoryImpl) GetAllTickets() ([]TicketResponse, error) {
	//----> Fetch all tickets.
	var tickets []models.Ticket
	if err := t.DB.Preload("Customer.User").Find(&tickets).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return ToTicketResponseList(tickets), nil
}

func (t *TicketRepositoryImpl) GetTicketByCustomerId(customerId string) ([]TicketResponse, error) {
	//----> Fetch all tickets.
	var tickets []models.Ticket
	if err := t.DB.Preload("Customer.User").Where("customer_id = ?", customerId).Find(&tickets).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return ToTicketResponseList(tickets), nil
}
