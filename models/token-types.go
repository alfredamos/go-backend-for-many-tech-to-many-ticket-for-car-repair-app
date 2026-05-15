package models

type TokenType string

const (
	Bearer TokenType = "Bearer"
)

type Status string

const (
	Valid   Status = "Valid"
	Invalid Status = "Invalid"
)

type QueryConditions struct {
	Status      Status
	UserID      string
	accessToken string
}

type TokenRequest struct {
	AccessToken  string    `json:"accessToken" validate:"required"`
	RefreshToken string    `json:"refreshToken" validate:"required"`
	TokenType    TokenType `json:"tokenType" validate:"required,oneof=Bearer"`
	Expired      bool      `json:"expired" default:"false"`
	Revoked      bool      `json:"revoked" default:"false"`
	TokenStatus  Status    `json:"status" validate:"required,oneof=Invalid Valid"`
	UserID       string    `json:"userId"`
}

type TokenResponse struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	TokenType    TokenType `json:"tokenType"`
	Expired      bool      `json:"expired"`
	Revoked      bool      `json:"revoked"`
	TokenStatus  Status    `json:"status"`
	UserID       string    `json:"userId"`
}

func fromTokenRequestToToken(request *TokenRequest) *Token {
	return &Token{
		AccessToken:  request.AccessToken,
		RefreshToken: request.RefreshToken,
		TokenType:    request.TokenType,
		Expired:      request.Expired,
		Revoked:      request.Revoked,
		UserID:       request.UserID,
		TokenStatus:  request.TokenStatus,
	}
}
