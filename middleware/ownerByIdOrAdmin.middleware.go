package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func OwnerByIdOrAdmin(ctx *fiber.Ctx) error {
	//----> Get user id from request.
	userId := ctx.Params("id")

	//----> Get session.
	session, err := GetSession(ctx)

	//----> Check for error in getting session.
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(errors.New(err.Error()))
	}

	//----> Check for owner or admin privilege.
	if !isOwner(userId, session.UserId) && !session.IsAdmin {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "You are not permitted to access this page"})
	}

	//----> Owner and admin.
	return ctx.Next()

}

func isOwner(userId, ownerId string) bool {
	return userId == ownerId
}
