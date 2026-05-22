package services

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"

	"github.com/gofiber/fiber/v2"
)

type AssignedTicketServiceInt interface {
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

type AssignedTicketServiceImpl struct {
	repo *repositories.AssignedTicketRepositoryImpl
}

func NewAssignedTicketServiceImpl(repo *repositories.AssignedTicketRepositoryImpl) *AssignedTicketServiceImpl {
	return &AssignedTicketServiceImpl{repo: repo}
}

func (a *AssignedTicketServiceImpl) ChangeAssignedTicketStatus(ctx *fiber.Ctx) error {
	//----> Get the assigned ticket id from payload.
	techId := ctx.Params("techId")
	ticketId := ctx.Params("ticketId")

	//----> Change assigned-ticket status.
	resp, err := a.repo.ChangeAssignedTicketStatus(techId, ticketId)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (a *AssignedTicketServiceImpl) CreateAssignedTicket(ctx *fiber.Ctx) error {
	//----> Initialize AssignedTicketCreate.
	assignedTicketCreate := models.AssignedTicketCreate{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&assignedTicketCreate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//----> Get user session.
	session, err := middleware.GetSession(ctx)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}

	//----> Input assign-by.
	assignedTicketCreate.AssignBy = session.Name

	resp, err := a.repo.CreateAssignedTicket(&assignedTicketCreate)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusCreated).JSON(resp)

}

func (a *AssignedTicketServiceImpl) DeleteAssignedTicketById(ctx *fiber.Ctx) error {
	//----> Get the assigned ticket id from payload.
	techId := ctx.Params("techId")
	ticketId := ctx.Params("ticketId")

	//----> Delete the assigned ticket with the giving id.
	resp, err := a.repo.DeleteAssignedTicketById(techId, ticketId)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (a *AssignedTicketServiceImpl) EditAssignedTicketById(ctx *fiber.Ctx) error {
	//----> Get the assigned ticket id from payload.
	techId := ctx.Params("techId")
	ticketId := ctx.Params("ticketId")

	//----> Initialize AssignedTicketEdit.
	assignedTicketEdit := models.AssignedTicketEdit{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&assignedTicketEdit); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//----> Get user session.
	session, err := middleware.GetSession(ctx)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}

	//----> Input assign-by.
	assignedTicketEdit.AssignBy = session.Name

	//----> Edit assigned-ticket.
	resp, err := a.repo.EditAssignedTicketId(techId, ticketId, &assignedTicketEdit)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (a *AssignedTicketServiceImpl) GetAssignedTicketById(ctx *fiber.Ctx) error {
	//----> Get the assigned ticket id from payload.
	techId := ctx.Params("techId")
	ticketId := ctx.Params("ticketId")

	//----> Fetch the assigned ticket with the giving id.
	resp, err := a.repo.GetAssignedTicketById(techId, ticketId)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (a *AssignedTicketServiceImpl) GetAllAssignedTickets(ctx *fiber.Ctx) error {
	//----> Fetch all assigned tickets.
	resp, err := a.repo.GetAllAssignedTickets()

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (a *AssignedTicketServiceImpl) GetAssignedTicketsByTechId(ctx *fiber.Ctx) error {
	//----> Get the assigned ticket id from payload.
	techId := ctx.Params("techId")

	//----> Fetch all assigned tickets.
	resp, err := a.repo.GetAssignedTicketsByTechId(techId)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (a *AssignedTicketServiceImpl) GetAssignedTicketsByTicketId(ctx *fiber.Ctx) error {
	//----> Get the assigned ticket id from payload.
	ticketId := ctx.Params("ticketId")

	//----> Fetch all assigned tickets.
	resp, err := a.repo.GetAssignedTicketsByTicketId(ticketId)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (a *AssignedTicketServiceImpl) GetCompletedAssignedTickets(ctx *fiber.Ctx) error {
	//----> Fetch all completed tickets.
	resp, err := a.repo.GetCompletedAssignedTicket()

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (a *AssignedTicketServiceImpl) GetInCompletedAssignedTickets(ctx *fiber.Ctx) error {
	//----> Fetch all in-completed tickets.
	resp, err := a.repo.GetInCompletedAssignedTicket()

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
