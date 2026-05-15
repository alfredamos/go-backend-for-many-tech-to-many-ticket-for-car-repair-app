package models

import (
	"time"

	"gorm.io/gorm"
)

type AssignedTicket struct {
	TechnicianID string `gorm:"primaryKey;type:varchar(255)" json:"userId"`
	TicketID     string `gorm:"primaryKey;type:varchar(255)" json:"eventId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	AssignAt     time.Time      `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"assignDate"`
	AssignBy     string         `json:"assignBy"`
	Completed    bool           `json:"completed" default:"false"`
	Status       TicketStatus   `json:"status" validate:"required,oneof=Closed Open" gorm:"default:'Open'"`
}
