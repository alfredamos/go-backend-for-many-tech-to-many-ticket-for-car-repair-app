package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string `json:"id" gorm:"primaryKey;type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `validate:"required" json:"name"`
	Password    string         `validate:"required" json:"password"`
	Phone       string         `validate:"required" json:"phone"`
	Email       string         `validate:"required,email" json:"email" gorm:"unique"`
	Role        Role           `validate:"required,role" json:"role"`
	Gender      Gender         `validate:"required,gender" json:"gender"`
	Image       string         `validate:"required,image" json:"image"`
	DateOfBirth time.Time      `json:"dateOfBirth" validate:"required"`

	Tokens     []Token    `json:"tokens" gorm:"foreignKey:UserID"`
	Customer   Customer   `json:"customer" gorm:"foreignKey:UserID"`
	Technician Technician `json:"technician" gorm:"foreignKey:UserID"`
}

// BeforeCreate These functions are called before creating any Post
func (user *User) BeforeCreate(_ *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
}
