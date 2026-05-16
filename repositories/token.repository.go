package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TokenRepositoryInt interface {
	CreateToken(token *models.TokenRequest) error
	DeleteInvalidTokensByUserId(userId string, ctx *fiber.Ctx) error
	DeleteAllInvalidTokens() error
	RevokeAllValidTokensByUserId(userId string) error
}

type TokenRepositoryImpl struct {
	DB *gorm.DB
}

func NewTokenRepositoryImpl(DB *gorm.DB) *TokenRepositoryImpl {
	return &TokenRepositoryImpl{DB: DB}
}

func (tResp TokenRepositoryImpl) CreateToken(token *models.TokenRequest) error {
	//----> Validate input.
	if err := models.ValidateTokenRequest(token); err != nil {
		return errors.New(err.Error())
	}

	//----> Map from token request to token model.
	tokenModel := fromTokenRequestToToken(token)

	//----> Create token.
	if err := tResp.DB.Create(tokenModel).Error; err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil
}

func (tResp TokenRepositoryImpl) DeleteInvalidTokensByUserId(userId string, ctx *fiber.Ctx) error {
	//----> Retrieve all invalid tokens.
	queryConditions := QueryConditions{UserID: userId, TokenStatus: models.Invalid}
	return deleteInvalidTokens(queryConditions, &tResp)
}

func (tResp TokenRepositoryImpl) DeleteAllInvalidTokens() error {
	//----> Retrieve all invalid tokens.
	queryConditions := QueryConditions{TokenStatus: models.Invalid}
	return deleteInvalidTokens(queryConditions, &tResp)
}

func (tResp TokenRepositoryImpl) RevokeAllValidTokensByUserId(userId string) error {
	//----> Retrieve all valid tokens.
	queryConditions := QueryConditions{UserID: userId, TokenStatus: models.Valid}
	tokens, err := findValidOrInvalidTokens(queryConditions, &tResp)

	//----> Check for errors in retrieving valid tokens.
	if err != nil {
		return errors.New(err.Error())
	}

	//----> Revoke all valid tokens
	if err := revokeValidTokens(tokens, &tResp); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back result.
	return nil

}
