package middleware

import (
	_ "encoding/json"
	"errors"
	_ "fmt"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

	"github.com/gofiber/fiber/v2"
)

type Session struct {
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	AccessToken string `json:"accessToken"`
	IsLoggedIn  bool   `json:"isLoggedIn"`
	IsAdmin     bool   `json:"isAdmin"`
}

func GetSession(ctx *fiber.Ctx) (Session, error) {
	//----> Get access-token from cookie.
	accessToken := GetCookie(ctx, "accessToken")

	//----> Validate the access-token on session.
	tokenJwt, err := ValidateToken(accessToken, ctx)

	//----> Check for error.
	if err != nil {
		return Session{}, errors.New(err.Error())
	}

	//----> Send back response
	return makeSession(tokenJwt, accessToken), nil
}

func makeSession(tokenJwt TokenJwt, accessToken string) Session {
	return Session{
		Email:       tokenJwt.Email,
		Name:        tokenJwt.Name,
		IsLoggedIn:  true,
		UserId:      tokenJwt.UserId,
		Role:        tokenJwt.Role,
		AccessToken: accessToken,
		IsAdmin:     tokenJwt.Role == string(models.AdminRole),
	}
}
