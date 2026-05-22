package controllers

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
)

type AssignedTicketControllerInt interface {
	ChangeAssignedTicketStatus(ctx *fiber.Ctx) error
	CreateAssignedTicket(ctx *fiber.Ctx) error
	DeleteAssignedTicketById(ctx *fiber.Ctx) error
	EditAssignedTicketById(ctx *fiber.Ctx) error
	GetAssignedTicketById(ctx *fiber.Ctx) error
	GetAllAssignedTickets(ctx *fiber.Ctx) error
	GetAssignedTicketsByTechId(ctx *fiber.Ctx) error
	GetAssignedTicketsByTicketId(ctx *fiber.Ctx) error
	GetCompletedAssignedTickets(ctx *fiber.Ctx) error
	GetInCompletedAssignedTickets(ctx *fiber.Ctx) error
}

type AssignedTicketControllerImpl struct {
	Service *services.AssignedTicketServiceImpl
}

func NewAssignedTicketControllerImpl(service *services.AssignedTicketServiceImpl) *AssignedTicketControllerImpl {
	return &AssignedTicketControllerImpl{Service: service}
}

func (a *AssignedTicketControllerImpl) ChangeAssignedTicketStatus(ctx *fiber.Ctx) error {
	return a.Service.ChangeAssignedTicketStatus(ctx)
}

func (a *AssignedTicketControllerImpl) CreateAssignedTicket(ctx *fiber.Ctx) error {
	return a.Service.CreateAssignedTicket(ctx)
}

func (a *AssignedTicketControllerImpl) DeleteAssignedTicketById(ctx *fiber.Ctx) error {
	return a.Service.DeleteAssignedTicketById(ctx)
}

func (a *AssignedTicketControllerImpl) EditAssignedTicketById(ctx *fiber.Ctx) error {
	return a.Service.EditAssignedTicketById(ctx)
}

func (a *AssignedTicketControllerImpl) GetAssignedTicketById(ctx *fiber.Ctx) error {
	return a.Service.GetAssignedTicketById(ctx)
}

func (a *AssignedTicketControllerImpl) GetAllAssignedTickets(ctx *fiber.Ctx) error {
	return a.Service.GetAllAssignedTickets(ctx)
}

func (a *AssignedTicketControllerImpl) GetAssignedTicketsByTechId(ctx *fiber.Ctx) error {
	return a.Service.GetAssignedTicketsByTechId(ctx)
}

func (a *AssignedTicketControllerImpl) GetAssignedTicketsByTicketId(ctx *fiber.Ctx) error {
	return a.Service.GetAssignedTicketsByTicketId(ctx)
}

func (a *AssignedTicketControllerImpl) GetCompletedAssignedTickets(ctx *fiber.Ctx) error {
	return a.Service.GetCompletedAssignedTickets(ctx)
}

func (a *AssignedTicketControllerImpl) GetInCompletedAssignedTickets(ctx *fiber.Ctx) error {
	return a.Service.GetInCompletedAssignedTickets(ctx)
}
