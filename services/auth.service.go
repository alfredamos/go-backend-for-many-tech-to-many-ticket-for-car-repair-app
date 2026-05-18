package services

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"

	"github.com/gofiber/fiber/v2"
)

type AuthServiceInt interface {
	ChangeUserPassword(ctx *fiber.Ctx) error
	ChangeUserRole(ctx *fiber.Ctx) error
	EditUserProfile(ctx *fiber.Ctx) error
	GetCurrentUser(ctx *fiber.Ctx) error
	GetUserSession(ctx *fiber.Ctx) error
	LoginUser(ctx *fiber.Ctx) error
	LogoutUser(ctx *fiber.Ctx) error
	RefreshUserToken(ctx *fiber.Ctx) error
	SignupUser(ctx *fiber.Ctx) error
}

type AuthServiceImpl struct {
	repo repositories.UserAuthRepoImpl
}

func NewAuthServiceImpl(repo repositories.UserAuthRepoImpl) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

func (a AuthServiceImpl) ChangeUserPassword(ctx *fiber.Ctx) error {
	//----> Initialize request object.
	request := models.ChangeUserPasswordRequest{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//----> Change user password.
	if err := a.repo.ChangeUserPassword(request); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON("password changed successfully")

}

func (a AuthServiceImpl) ChangeUserRole(ctx *fiber.Ctx) error {
	//----> Initialize request object.
	request := models.ChangeUserRoleRequest{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//----> Change user role.
	userResp, err := a.repo.ChangeUserRole(ctx, request)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(userResp)
}

func (a AuthServiceImpl) EditUserProfile(ctx *fiber.Ctx) error {
	//----> Initialize request object.
	request := models.EditUserProfileRequest{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//----> Edit user profile.
	if err := a.repo.EditUserProfile(request); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON("profile updated successfully")

}

func (a AuthServiceImpl) GetCurrentUser(ctx *fiber.Ctx) error {
	//----> Get the current user.
	userResp, err := a.repo.GetCurrentUser(ctx)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(userResp)
}

func (a AuthServiceImpl) GetUserSession(ctx *fiber.Ctx) error {
	//----> Get the user session.
	session, err := a.repo.GetUserSession(ctx)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(session)
}

func (a AuthServiceImpl) LoginUser(ctx *fiber.Ctx) error {
	//----> Initialize request object.
	request := models.LoginUserRequest{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//----> Login user.
	session, err := a.repo.LoginUser(ctx, request)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(session)
}

func (a AuthServiceImpl) LogoutUser(ctx *fiber.Ctx) error {
	//----> Logout user.
	session, err := a.repo.LogoutUser(ctx)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(session)
}

func (a AuthServiceImpl) RefreshUserToken(ctx *fiber.Ctx) error {
	//----> Refresh user token.
	session, err := a.repo.RefreshUserToken(ctx)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(session)
}

func (a AuthServiceImpl) SignupUser(ctx *fiber.Ctx) error {
	//----> Initialize request object.
	request := models.SignupUserRequest{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//----> Signup user.
	if err := a.repo.SignupUser(request); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON("user created successfully")
}
