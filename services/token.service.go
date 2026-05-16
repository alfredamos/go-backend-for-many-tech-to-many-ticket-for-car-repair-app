package services

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"

	"github.com/gofiber/fiber/v2"
)

type TokenService interface {
	DeleteInvalidTokensByUserIdController(ctx *fiber.Ctx) error
	DeleteAllInvalidTokensController(ctx *fiber.Ctx) error
}

type TokenServiceImpl struct {
	repo repositories.TokenRepositoryImpl
}

func NewTokenServiceImpl(repo repositories.TokenRepositoryImpl) *TokenServiceImpl {
	return &TokenServiceImpl{repo: repo}
}

func (tkr TokenServiceImpl) DeleteInvalidTokensByUserIdController(ctx *fiber.Ctx) error {
	//----> Get the user id from payload.
	userId := ctx.Params("userId")

	//----> Delete all invalid tokens.
	if err := tkr.repo.DeleteInvalidTokensByUserId(userId, ctx); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON("tokens deleted successfully")

}

func (tkr TokenServiceImpl) DeleteAllInvalidTokensController(ctx *fiber.Ctx) error {
	//----> Delete all invalid tokens.
	if err := tkr.repo.DeleteAllInvalidTokens(); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON("tokens deleted successfully")
}
