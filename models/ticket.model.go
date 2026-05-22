package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID          string `json:"id" gorm:"primaryKey;type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string `json:"title" validate:"required"`
	Description string `json:"notes" validate:"required"`

	Technicians []Technician `gorm:"many2many:assignedTickets;"`

	CustomerID string    `json:"customerId"`
	Customer   *Customer `json:"customer" gorm:"foreignKey:CustomerID"`
}

// BeforeCreate These functions are called before creating any Post
func (ticket *Ticket) BeforeCreate(_ *gorm.DB) (err error) {
	ticket.ID = uuid.New().String()
	return
}
