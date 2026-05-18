package controllers

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
)

type TicketControllerInt interface {
	CreateTicketController(ctx *fiber.Ctx) error
	DeleteTicketByIdController(ctx *fiber.Ctx) error
	EditTicketByIdController(ctx *fiber.Ctx) error
	GetAllTicketsController(ctx *fiber.Ctx) error
	GetTicketByIdController(ctx *fiber.Ctx) error
	GetTicketsByCustomerIdController(ctx *fiber.Ctx) error
}

type TicketControllerImpl struct {
	Service *services.TicketServiceImpl
}

func NewTicketControllerImpl(service *services.TicketServiceImpl) *TicketControllerImpl {
	return &TicketControllerImpl{Service: service}
}

func (t *TicketControllerImpl) CreateTicketController(ctx *fiber.Ctx) error {
	return t.Service.CreateTicket(ctx)
}

func (t *TicketControllerImpl) DeleteTicketByIdController(ctx *fiber.Ctx) error {
	return t.Service.DeleteTicketById(ctx)
}

func (t *TicketControllerImpl) EditTicketByIdController(ctx *fiber.Ctx) error {
	return t.Service.EditTicketById(ctx)
}

func (t *TicketControllerImpl) GetAllTicketsController(ctx *fiber.Ctx) error {
	return t.Service.GetAllTickets(ctx)
}

func (t *TicketControllerImpl) GetTicketByIdController(ctx *fiber.Ctx) error {
	return t.Service.GetTicketById(ctx)
}

func (t *TicketControllerImpl) GetTicketsByCustomerIdController(ctx *fiber.Ctx) error {
	return t.Service.GetTicketsByCustomerId(ctx)
}
