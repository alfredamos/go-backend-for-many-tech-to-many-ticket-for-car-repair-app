package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func OwnerCheckOrAdminByUserId(ctx *fiber.Ctx, userId string) (bool, error) {
	//----> Get session.
	session, err := GetSession(ctx)

	//----> Check for error in getting session.
	if err != nil {
		return false, errors.New(err.Error())
	}

	//----> Check for admin privilege.
	isAdmin := session.IsAdmin

	//----> Check for ownership.
	isOwner := session.UserId == userId

	//----> Not owner and not admin.
	if !isAdmin && !isOwner {
		return false, errors.New("you are not allowed to view or perform this action")
	}

	//----> Admin or owner.
	return true, nil
}

func OwnerCheckOrAdminByEmail(ctx *fiber.Ctx, email string) (bool, error) {
	//----> Get session.
	session, err := GetSession(ctx)

	//----> Check for error in getting session.
	if err != nil {
		return false, errors.New(err.Error())
	}

	//----> Check for admin privilege.
	isAdmin := session.IsAdmin

	//----> Check for ownership.
	isOwner := session.Email == email

	//----> Not owner and not admin.
	if !isAdmin && !isOwner {
		return false, errors.New("you are not allowed to view or perform this action")
	}

	//----> Admin or owner.
	return isAdmin || isOwner, nil
}

func IsAdmin(ctx *fiber.Ctx) (bool, error) {
	//----> Get session.
	session, err := GetSession(ctx)

	//----> Check for error in getting session.
	if err != nil {
		return false, errors.New(err.Error())
	}

	//----> Not owner and not admin.
	if !session.IsAdmin {
		return false, errors.New("you are not allowed to perform this action")
	}

	//----> Must be admin.
	return session.IsAdmin, nil
}
