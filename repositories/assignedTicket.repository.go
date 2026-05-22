package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

	"gorm.io/gorm"
)

type AssignedTicketRepositoryInt interface {
	ChangeAssignedTicketStatus(techId, ticketId string) (ResponseMessage, error)
	CreateAssignedTicket(request *models.AssignedTicketCreate) (ResponseMessage, error)
	DeleteAssignedTicketById(techId, ticketId string) (ResponseMessage, error)
	EditAssignedTicketById(techId, ticketId string, request *models.AssignedTicketEdit) (ResponseMessage, error)
	GetAssignedTicketById(techId, ticketId string) (AssignedTicketResponse, error)
	GetAllAssignedTickets() ([]AssignedTicketResponse, error)
	GetAssignedTicketsByTechId(techId string) ([]AssignedTicketResponse, error)
	GetAssignedTicketsByTicketId(ticketId string) ([]AssignedTicketResponse, error)
	GetCompletedAssignedTicket() ([]AssignedTicketResponse, error)
	GetInCompletedAssignedTicket() ([]AssignedTicketResponse, error)
}

type AssignedTicketRepositoryImpl struct {
	DB *gorm.DB
}

func NewAssignedTicketRepositoryImpl(DB *gorm.DB) *AssignedTicketRepositoryImpl {
	return &AssignedTicketRepositoryImpl{DB: DB}
}

func (a *AssignedTicketRepositoryImpl) ChangeAssignedTicketStatus(techId, ticketId string) (ResponseMessage, error) {
	//----> Get the assigned ticket.
	ticket, err := getOneAssignedTicketById(techId, ticketId, a)

	//----> Check for error.
	if err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Initialize assigned ticket.
	ticket.Completed = !ticket.Completed
	ticket.Status = changeStatus(ticket.Completed)
	assignedTicket := models.AssignedTicket{
		TicketID:     ticketId,
		TechnicianID: techId,
		Completed:    ticket.Completed,
		Status:       ticket.Status,
		AssignAt:     ticket.AssignAt,
	}

	//----> Update the assigned ticket status.
	assignedTicket.CreatedAt = ticket.CreatedAt
	if err := a.DB.Save(&assignedTicket).Error; err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Return the updated assigned ticket.
	return NewResponseMessage("Assigned ticket status changed successfully", 200, "Success"), nil
}

func (a *AssignedTicketRepositoryImpl) CreateAssignedTicket(request *models.AssignedTicketCreate) (ResponseMessage, error) {
	//----> Validate input.
	if err := models.ValidateAssignedTicketCreate(request); err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Create assigned ticket.
	assignedTicket := models.AssignedTicket{
		TechnicianID: request.TechnicianID,
		TicketID:     request.TicketID,
		AssignBy:     request.AssignBy,
	}

	//----> Save assigned ticket.
	if err := a.DB.Create(&assignedTicket).Error; err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Return the created assigned ticket.
	return NewResponseMessage("Assigned ticket created successfully", 201, "Success"), nil
}

func (a *AssignedTicketRepositoryImpl) DeleteAssignedTicketById(techId, ticketId string) (ResponseMessage, error) {
	//----> Check for existence of assigned-ticket.
	ticket, err := getOneAssignedTicketById(techId, ticketId, a)

	//----> Check for error.
	if err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Delete assigned-ticket.
	if err := a.DB.Delete(&ticket).Error; err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Return the deleted assigned-ticket.
	return NewResponseMessage("Assigned ticket deleted successfully", 200, "Success"), nil
}

func (a *AssignedTicketRepositoryImpl) EditAssignedTicketId(techId, ticketId string, request *models.AssignedTicketEdit) (ResponseMessage, error) {
	//----> Check for existence of assigned-ticket.
	_, err := getOneAssignedTicketById(techId, ticketId, a)

	//----> Check for error.
	if err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Validate input.
	if err := models.ValidateAssignedTicketEdit(request); err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Initialize assigned-ticket.
	ticket := models.AssignedTicket{
		TechnicianID: request.TechnicianID,
		TicketID:     request.TicketID,
		AssignBy:     request.AssignBy,
		Completed:    request.Completed,
		Status:       request.Status,
	}

	//----> Update assigned-ticket.
	if err := a.DB.Updates(ticket).Error; err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Return the updated assigned-ticket.
	return NewResponseMessage("Assigned ticket updated successfully", 200, "Success"), nil
}

func (a *AssignedTicketRepositoryImpl) GetAssignedTicketById(techId, ticketId string) (AssignedTicketResponse, error) {
	//----> Fetch the assigned-ticket with the giving id.
	ticket, err := getOneAssignedTicketById(techId, ticketId, a)

	//----> Check for error.
	if err != nil {
		return AssignedTicketResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toAssignedTicketResponse(ticket), nil
}

func (a *AssignedTicketRepositoryImpl) GetAllAssignedTickets() ([]AssignedTicketResponse, error) {
	//----> Initialize assigned-tickets.
	var assignedTickets []models.AssignedTicket

	//----> Fetch all assigned-tickets.
	if err := a.DB.Preload("Ticket.Customer.User").Preload("Technician.User").Find(&assignedTickets).Error; err != nil {
		return []AssignedTicketResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toAssignedTicketResponseList(assignedTickets), nil
}

func (a *AssignedTicketRepositoryImpl) GetAssignedTicketsByTechId(techId string) ([]AssignedTicketResponse, error) {
	//----> Initialize assigned-tickets.
	var assignedTickets []models.AssignedTicket

	//----> Fetch all assigned-tickets.
	if err := a.DB.Preload("Ticket.Customer.User").Preload("Technician.User").Where("Technician_id = ?", techId).Find(&assignedTickets).Error; err != nil {
		return []AssignedTicketResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toAssignedTicketResponseList(assignedTickets), nil
}

func (a *AssignedTicketRepositoryImpl) GetAssignedTicketsByTicketId(ticketId string) ([]AssignedTicketResponse, error) {
	//----> Initialize assigned-tickets.
	var assignedTickets []models.AssignedTicket

	//----> Fetch all assigned-tickets.
	if err := a.DB.Preload("Ticket.Customer.User").Preload("Technician.User").Where("ticket_id = ?", ticketId).Find(&assignedTickets).Error; err != nil {
		return []AssignedTicketResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toAssignedTicketResponseList(assignedTickets), nil
}

func (a *AssignedTicketRepositoryImpl) GetCompletedAssignedTicket() ([]AssignedTicketResponse, error) {
	//----> Initialize assigned-tickets.
	var assignedTickets []models.AssignedTicket

	//----> Fetch all assigned-tickets.
	if err := a.DB.Preload("Ticket.Customer.User").Preload("Technician.User").Where("completed = ?", true).Find(&assignedTickets).Error; err != nil {
		return []AssignedTicketResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toAssignedTicketResponseList(assignedTickets), nil
}

func (a *AssignedTicketRepositoryImpl) GetInCompletedAssignedTicket() ([]AssignedTicketResponse, error) {
	//----> Initialize assigned-tickets.
	var assignedTickets []models.AssignedTicket

	//----> Fetch all assigned-tickets.
	if err := a.DB.Preload("Ticket.Customer.User").Preload("Technician.User").Where("completed = ?", false).Find(&assignedTickets).Error; err != nil {
		return []AssignedTicketResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toAssignedTicketResponseList(assignedTickets), nil
}

func changeStatus(completed bool) models.TicketStatus {
	Status := models.TicketStatus("Open")
	if completed {
		Status = "Closed"
	}
	return Status
}
