package repositories

import "go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

type TokenRequest struct {
	AccessToken  string           `json:"accessToken" validate:"required"`
	RefreshToken string           `json:"refreshToken" validate:"required"`
	TokenType    models.TokenType `json:"tokenType" validate:"required,oneof=Bearer"`
	Expired      bool             `json:"expired" default:"false"`
	Revoked      bool             `json:"revoked" default:"false"`
	TokenStatus  models.Status    `json:"status" validate:"required,oneof=Invalid Valid"`
	UserID       string           `json:"userId"`
}

type TokenResponse struct {
	AccessToken  string           `json:"accessToken"`
	RefreshToken string           `json:"refreshToken"`
	TokenType    models.TokenType `json:"tokenType"`
	Expired      bool             `json:"expired"`
	Revoked      bool             `json:"revoked"`
	TokenStatus  models.Status    `json:"status"`
	UserID       string           `json:"userId"`
}
