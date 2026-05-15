package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type TokenJwt struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	UserId string `json:"userId"`
}

func GetUserIdEmailNameRole(c *fiber.Ctx) TokenJwt {
	//----> Get user name
	email := fmt.Sprintf("%v", c.Locals("email"))

	//----> Get user name
	name := fmt.Sprintf("%v", c.Locals("name"))

	//----> Get user role from context.
	role := fmt.Sprintf("%v", c.Locals("role"))

	//----> Get the user-id from context.
	userId := fmt.Sprintf("%v", c.Locals("userId"))

	//----> Send back user-detail
	return TokenJwt{Email: email, Name: name, Role: role, UserId: userId}
}

func GetUserIdFromContext(c *fiber.Ctx) string {
	//----> Get user-id from context.
	userId := c.Locals("userId")

	//----> Send back the user-id.
	return fmt.Sprintf("%v", userId)
}
