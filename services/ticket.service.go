package services

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"

	"github.com/gofiber/fiber/v2"
)

type TicketServiceInt interface {
	CreateTicket(ctx *fiber.Ctx) error
	DeleteTicketById(ctx *fiber.Ctx) error
	EditTicketById(ctx *fiber.Ctx) error
	GetAllTickets(ctx *fiber.Ctx) error
	GetTicketById(ctx *fiber.Ctx) error
	GetTicketsByCustomerId(ctx *fiber.Ctx) error
}

type TicketServiceImpl struct {
	repo *repositories.TicketRepositoryImpl
}

func NewTicketServiceImpl(repo *repositories.TicketRepositoryImpl) *TicketServiceImpl {
	return &TicketServiceImpl{repo: repo}
}

func (t *TicketServiceImpl) CreateTicket(ctx *fiber.Ctx) error {
	//----> Initialize ticket create object.
	ticketCreate := models.TicketCreate{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&ticketCreate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//----> Create ticket.
	ticket, err := t.repo.CreateTicket(&ticketCreate)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusCreated).JSON(ticket)
}

func (t *TicketServiceImpl) DeleteTicketById(ctx *fiber.Ctx) error {
	//----> Get the ticket id from payload.
	id := ctx.Params("id")

	//----> Delete the ticket with the giving id.
	ticket, err := t.repo.DeleteTicketById(id)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(ticket)
}

func (t *TicketServiceImpl) EditTicketById(ctx *fiber.Ctx) error {
	//----> Get the ticket id from payload.
	id := ctx.Params("id")

	//----> Initialize ticket edit object.
	ticketEdit := models.TicketEdit{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&ticketEdit); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//----> Edit the ticket with the giving id.
	ticket, err := t.repo.EditTicketById(id, &ticketEdit)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(ticket)
}

func (t *TicketServiceImpl) GetTicketById(ctx *fiber.Ctx) error {
	//----> Get the ticket id from payload.
	id := ctx.Params("id")

	//----> Fetch the ticket with the giving id.
	ticket, err := t.repo.GetTicketById(id)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(ticket)
}

func (t *TicketServiceImpl) GetAllTickets(ctx *fiber.Ctx) error {
	//----> Fetch all tickets.
	tickets, err := t.repo.GetAllTickets()

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(tickets)
}

func (t *TicketServiceImpl) GetTicketsByCustomerId(ctx *fiber.Ctx) error {
	//----> Get the customer id from payload.
	customerId := ctx.Params("customerId")

	//----> Fetch tickets by customer id.
	tickets, err := t.repo.GetTicketByCustomerId(customerId)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(tickets)
}
