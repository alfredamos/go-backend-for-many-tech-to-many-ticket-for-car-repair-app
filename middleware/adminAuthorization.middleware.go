package middleware

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AdminRolePermission(ctx *fiber.Ctx) error {
	//----> Get session.
	session, err := GetSession(ctx)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(errors.New(err.Error()))
	}

	//----> Check for admin privilege.
	if !session.IsAdmin {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "You are not permitted to access this page", "statusCode": http.StatusForbidden})
	}

	//----> You are permitted.
	return ctx.Next()
}
