package controllers

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
)

type UserControllerInt interface {
	DeleteUserByIdController(ctx *fiber.Ctx) error
	GetUserByIdController(ctx *fiber.Ctx) error
	GetUserByEmailController(ctx *fiber.Ctx) error
	GetAllUsersController(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	service services.UserServiceImpl
}

func NewUserControllerImpl(service services.UserServiceImpl) *UserControllerImpl {
	return &UserControllerImpl{service: service}
}

func (u *UserControllerImpl) DeleteUserByIdController(ctx *fiber.Ctx) error {
	return u.service.DeleteUserById(ctx)
}

func (u *UserControllerImpl) GetUserByIdController(ctx *fiber.Ctx) error {
	return u.service.GetUserById(ctx)
}

func (u *UserControllerImpl) GetUserByEmailController(ctx *fiber.Ctx) error {
	return u.service.GetUserByEmail(ctx)
}

func (u *UserControllerImpl) GetAllUsersController(ctx *fiber.Ctx) error {
	return u.service.GetAllUsers(ctx)
}
