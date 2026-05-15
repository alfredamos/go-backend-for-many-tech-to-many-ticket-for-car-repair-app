package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID        string         `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Address   string         `json:"address" validate:"required"`
	Notes     string         `json:"notes" validate:"required"`
	Active    bool           `json:"active" default:"true"`
	Status    Status         `json:"status" default:"Invalid"`

	UserID  string   `json:"userId" gorm:"unique;not null"`
	Tickets []Ticket `json:"tickets" gorm:"foreignKey:CustomerID"`
}

// BeforeCreate These functions are called before creating any Post
func (customer *Customer) BeforeCreate(_ *gorm.DB) (err error) {
	customer.ID = uuid.New().String()
	return
}
