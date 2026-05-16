package services

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"

	"github.com/gofiber/fiber/v2"
)

type UserServiceInt interface {
	DeleteUserById(ctx *fiber.Ctx) error
	GetUserById(ctx *fiber.Ctx) error
	GetUserByEmail(ctx *fiber.Ctx) error
	GetAllUsers(ctx *fiber.Ctx) error
}

type UserServiceImpl struct {
	service repositories.UserRepositoryImpl
}

func NewUserServiceImpl(service repositories.UserRepositoryImpl) *UserServiceImpl {
	return &UserServiceImpl{service: service}
}

func (u *UserServiceImpl) DeleteUserById(ctx *fiber.Ctx) error {
	//----> Get the user id from payload.
	id := ctx.Params("id")

	//----> Delete the user with the giving id.
	user, err := u.service.DeleteUserById(id)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (u *UserServiceImpl) GetUserById(ctx *fiber.Ctx) error {
	//----> Get the user id from payload.
	id := ctx.Params("id")

	//----> Fetch the user with the giving id.
	user, err := u.service.GetUserById(id)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (u *UserServiceImpl) GetUserByEmail(ctx *fiber.Ctx) error {
	//----> Get the user email from payload.
	email := ctx.Params("email")

	//----> Fetch the user with the giving email.
	user, err := u.service.GetUserByEmail(email)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (u *UserServiceImpl) GetAllUsers(ctx *fiber.Ctx) error {
	//----> Fetch all users.
	users, err := u.service.GetAllUsers()

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(users)
}
