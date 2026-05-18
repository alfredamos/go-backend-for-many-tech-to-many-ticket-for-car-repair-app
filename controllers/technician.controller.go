package controllers

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
)

type TechnicianControllerInt interface {
	CreateTechnicianController(ctx *fiber.Ctx) error
	DeleteTechnicianByIdController(ctx *fiber.Ctx) error
	EditTechnicianByIdController(ctx *fiber.Ctx) error
	GetTechnicianByIdController(ctx *fiber.Ctx) error
	GetAllTechniciansController(ctx *fiber.Ctx) error
	GetTechnicianByUserIdController(ctx *fiber.Ctx) error
	GetTechnicianBySpecialtyController(ctx *fiber.Ctx) error
}

type TechnicianControllerImpl struct {
	controller *services.TechnicianServiceImpl
}

func NewTechnicianControllerImpl(controller *services.TechnicianServiceImpl) *TechnicianControllerImpl {
	return &TechnicianControllerImpl{controller: controller}
}

func (t *TechnicianControllerImpl) CreateTechnicianController(ctx *fiber.Ctx) error {
	return t.controller.CreateTechnician(ctx)
}

func (t *TechnicianControllerImpl) DeleteTechnicianByIdController(ctx *fiber.Ctx) error {
	return t.controller.DeleteTechnicianById(ctx)
}

func (t *TechnicianControllerImpl) EditTechnicianByIdController(ctx *fiber.Ctx) error {
	return t.controller.EditTechnicianById(ctx)
}

func (t *TechnicianControllerImpl) GetTechnicianByIdController(ctx *fiber.Ctx) error {
	return t.controller.GetTechnicianById(ctx)
}

func (t *TechnicianControllerImpl) GetAllTechniciansController(ctx *fiber.Ctx) error {
	return t.controller.GetAllTechnicians(ctx)
}

func (t *TechnicianControllerImpl) GetTechniciansBySpecialtyController(ctx *fiber.Ctx) error {
	return t.controller.GetTechnicianBySpecialty(ctx)
}

func (t *TechnicianControllerImpl) GetTechnicianByUserIdController(ctx *fiber.Ctx) error {
	return t.controller.GetTechnicianByUserId(ctx)
}
