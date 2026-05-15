package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	ID           string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	AccessToken  string         `json:"accessToken" gorm:"unique;type:varchar(750)" validate:"required"`
	RefreshToken string         `json:"refreshToken" gorm:"unique;type:varchar(750)" validate:"required"`
	TokenType    TokenType      `json:"tokenType" validate:"required"`
	Expired      bool           `json:"expired" default:"false"`
	Revoked      bool           `json:"revoked" default:"false"`
	TokenStatus  Status         `json:"status" validate:"required"`
	UserID       string         `json:"userId"`
}

// BeforeCreate These functions are called before creating any Post
func (token *Token) BeforeCreate(_ *gorm.DB) (err error) {
	token.ID = uuid.New().String()
	return
}
