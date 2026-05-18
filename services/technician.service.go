package services

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"

	"github.com/gofiber/fiber/v2"
)

type TechnicianServiceInt interface {
	CreateTechnician(ctx *fiber.Ctx) error
	DeleteTechnicianById(ctx *fiber.Ctx) error
	EditTechnicianById(ctx *fiber.Ctx) error
	GetTechnicianById(ctx *fiber.Ctx) error
	GetAllTechnicians(ctx *fiber.Ctx) error
	GetTechnicianByUserId(ctx *fiber.Ctx) error
	GetTechnicianBySpecialty(ctx *fiber.Ctx) error
}

type TechnicianServiceImpl struct {
	service repositories.TechnicianRepositoryImpl
}

func NewTechnicianServiceImpl(service repositories.TechnicianRepositoryImpl) *TechnicianServiceImpl {
	return &TechnicianServiceImpl{service: service}
}

func (t *TechnicianServiceImpl) CreateTechnician(ctx *fiber.Ctx) error {
	//----> Initialize technician create object.
	technicianCreate := models.TechnicianCreate{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&technicianCreate); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Create technician.
	tech, err := t.service.CreateTechnician(&technicianCreate)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusCreated).JSON(tech)
}

func (t *TechnicianServiceImpl) DeleteTechnicianById(ctx *fiber.Ctx) error {
	//----> Get the technician id from payload.
	id := ctx.Params("id")

	//----> Delete the technician with the giving id.
	tech, err := t.service.DeleteTechnicianById(id)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(tech)
}

func (t *TechnicianServiceImpl) EditTechnicianById(ctx *fiber.Ctx) error {
	//----> Get the technician id from payload.
	id := ctx.Params("id")

	//----> Initialize technician edit object.
	technicianEdit := models.TechnicianEdit{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&technicianEdit); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Edit the technician with the giving id.
	tech, err := t.service.EditTechnicianById(id, &technicianEdit)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(tech)
}

func (t *TechnicianServiceImpl) GetTechnicianById(ctx *fiber.Ctx) error {
	//----> Get the technician id from payload.
	id := ctx.Params("id")

	//----> Fetch the technician with the giving id.
	tech, err := t.service.GetTechnicianById(id)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(tech)
}

func (t *TechnicianServiceImpl) GetAllTechnicians(ctx *fiber.Ctx) error {
	//----> Fetch all technicians.
	technicians, err := t.service.GetAllTechnicians()

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(technicians)
}

func (t *TechnicianServiceImpl) GetTechnicianBySpecialty(ctx *fiber.Ctx) error {
	//----> Get the specialty from payload.
	specialty := ctx.Params("specialty")

	//----> Fetch the technician with the giving specialty.
	technicians, err := t.service.GetTechnicianBySpecialty(specialty)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(technicians)
}

func (t *TechnicianServiceImpl) GetTechnicianByUserId(ctx *fiber.Ctx) error {
	//----> Get the user id from payload.
	userId := ctx.Params("userId")

	//----> Fetch the technician with the giving user id.
	technician, err := t.service.GetTechnicianByUserId(userId)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(technician)
}
