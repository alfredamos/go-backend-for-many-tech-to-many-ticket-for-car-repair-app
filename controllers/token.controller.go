package controllers

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
)

type TokenControllerInt interface {
	DeleteInvalidTokensByUserIdController(ctx *fiber.Ctx) error
	DeleteAllInvalidTokensController(ctx *fiber.Ctx) error
}

type TokenControllerImpl struct {
	service *services.TokenServiceImpl
}

func NewTokenControllerImpl(service *services.TokenServiceImpl) *TokenControllerImpl {
	return &TokenControllerImpl{service: service}
}

func (tcl TokenControllerImpl) DeleteInvalidTokensByUserIdController(ctx *fiber.Ctx) error {
	//----> Delete invalid tokens by user id.
	return tcl.service.DeleteInvalidTokensByUserIdController(ctx)
}

func (tcl TokenControllerImpl) DeleteAllInvalidTokensController(ctx *fiber.Ctx) error {
	//----> Delete all invalid tokens.
	return tcl.service.DeleteAllInvalidTokensController(ctx)
}
