package controllers

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerInt interface {
	ChangeUserPasswordController(ctx *fiber.Ctx) error
	ChangeUserRoleController(ctx *fiber.Ctx) error
	EditUserProfileController(ctx *fiber.Ctx) error
	GetCurrentUserController(ctx *fiber.Ctx) error
	GetUserSessionController(ctx *fiber.Ctx) error
	LoginUserController(ctx *fiber.Ctx) error
	LogoutUserController(ctx *fiber.Ctx) error
	RefreshUserTokenController(ctx *fiber.Ctx) error
	SignupUserController(ctx *fiber.Ctx) error
}

type AuthControllerImpl struct {
	service services.AuthServiceImpl
}

func NewAuthController(service services.AuthServiceImpl) *AuthControllerImpl {
	return &AuthControllerImpl{service: service}
}

func (a AuthControllerImpl) ChangeUserPasswordController(ctx *fiber.Ctx) error {
	//----> Change user password.
	return a.service.ChangeUserPassword(ctx)
}

func (a AuthControllerImpl) ChangeUserRoleController(ctx *fiber.Ctx) error {
	//----> Change user role.
	return a.service.ChangeUserRole(ctx)
}

func (a AuthControllerImpl) EditUserProfileController(ctx *fiber.Ctx) error {
	//----> Edit user profile.
	return a.service.EditUserProfile(ctx)
}

func (a AuthControllerImpl) GetCurrentUserController(ctx *fiber.Ctx) error {
	//----> Get current user.
	return a.service.GetCurrentUser(ctx)
}

func (a AuthControllerImpl) GetUserSessionController(ctx *fiber.Ctx) error {
	//----> Get user session.
	return a.service.GetUserSession(ctx)
}

func (a AuthControllerImpl) LoginUserController(ctx *fiber.Ctx) error {
	//----> Login user.
	return a.service.LoginUser(ctx)
}

func (a AuthControllerImpl) LogoutUserController(ctx *fiber.Ctx) error {
	//----> Logout user.
	return a.service.LogoutUser(ctx)
}

func (a AuthControllerImpl) RefreshUserTokenController(ctx *fiber.Ctx) error {
	//----> Refresh user token.
	return a.service.RefreshUserToken(ctx)
}

func (a AuthControllerImpl) SignupUserController(ctx *fiber.Ctx) error {
	//----> Signup user.
	return a.service.SignupUser(ctx)
}
