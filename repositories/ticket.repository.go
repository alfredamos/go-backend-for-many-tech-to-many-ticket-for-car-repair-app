package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

	"gorm.io/gorm"
)

type TicketRepositoryInt interface {
	CreateTicket(ticketCreate *models.TicketCreate) (TicketResponse, error)
	DeleteTicketById(id string) (TicketResponse, error)
	EditTicketById(id string, ticketEdit *models.TicketEdit) (TicketResponse, error)
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

func (t *TicketRepositoryImpl) CreateTicket(ticketCreate *models.TicketCreate) (TicketResponse, error) {
	//----> Validate ticket input.
	if err := models.ValidateTicketCreate(ticketCreate); err != nil {
		return TicketResponse{}, err
	}

	//----> Initialize ticket.
	ticket := &models.Ticket{
		Title:       ticketCreate.Title,
		Description: ticketCreate.Description,
		CustomerID:  ticketCreate.CustomerID,
	}

	//----> Create ticket.
	if err := t.DB.Create(&ticket).Error; err != nil {
		return TicketResponse{}, err
	}

	//----> Initialize customer and user
	customer := &models.Customer{}
	user := &models.User{}

	//----> Preload customer and user.
	if err := t.DB.Where("id = ?", ticket.CustomerID).First(customer).Error; err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	if err := t.DB.Where("id = ?", customer.UserID).First(user).Error; err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toTicketResponse(*ticket, user), nil

}

func (t *TicketRepositoryImpl) DeleteTicketById(id string) (TicketResponse, error) {
	//----> Check for existence of ticket.
	ticket, err := getOneTicketById(id, t)

	//----> Check for error.
	if err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	//----> Delete ticket.
	if err := t.DB.Delete(&ticket).Error; err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	//----> Initialize customer and user
	customer := &models.Customer{}
	user := &models.User{}

	//----> Preload customer and user.
	if err := t.DB.Where("id = ?", ticket.CustomerID).First(customer).Error; err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	if err := t.DB.Where("id = ?", customer.UserID).First(user).Error; err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toTicketResponse(*ticket, user), nil
}

func (t *TicketRepositoryImpl) EditTicketById(id string, ticketEdit *models.TicketEdit) (TicketResponse, error) {
	//----> Check for existence of ticket.
	ticket, err := getOneTicketById(id, t)

	//----> Check for error.
	if err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	//----> Validate ticket input.
	if err := models.ValidateTicketEdit(ticketEdit); err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	//----> Initialize ticket.
	ticketToEdit := &models.TicketEdit{
		ID:          ticketEdit.ID,
		Title:       ticketEdit.Title,
		Description: ticketEdit.Description,
		CustomerID:  ticketEdit.CustomerID,
	}

	//----> Update ticket.
	if err := t.DB.Model(ticketEdit).Where("id = ?", id).Updates(ticketToEdit).Error; err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	//----> Initialize customer and user
	customer := &models.Customer{}
	user := &models.User{}

	//----> Preload customer and user.
	if err := t.DB.Where("id = ?", ticket.CustomerID).First(customer).Error; err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	if err := t.DB.Where("id = ?", customer.UserID).First(user).Error; err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toTicketResponse(*ticket, user), nil
}

func (t *TicketRepositoryImpl) GetTicketById(id string) (TicketResponse, error) {
	//----> Fetch the ticket with the giving id.
	ticket, err := getOneTicketById(id, t)

	//----> Check for error.
	if err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	//----> Initialize customer and user
	customer := &models.Customer{}
	user := &models.User{}

	//----> Preload customer and user.
	if err := t.DB.Where("id = ?", ticket.CustomerID).First(customer).Error; err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	if err := t.DB.Where("id = ?", customer.UserID).First(user).Error; err != nil {
		return TicketResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toTicketResponse(*ticket, user), nil

}

func (t *TicketRepositoryImpl) GetAllTickets() ([]TicketResponse, error) {
	//----> Fetch all tickets.
	var tickets []models.Ticket
	if err := t.DB.Find(&tickets).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return ToTicketResponseList(tickets, t), nil
}

func (t *TicketRepositoryImpl) GetTicketByCustomerId(customerId string) ([]TicketResponse, error) {
	//----> Fetch all tickets.
	var tickets []models.Ticket
	if err := t.DB.Where("customer_id = ?", customerId).Find(&tickets).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return ToTicketResponseList(tickets, t), nil
}
