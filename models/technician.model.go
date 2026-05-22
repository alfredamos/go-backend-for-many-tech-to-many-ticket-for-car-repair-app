package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Technician struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(255)"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    string         `json:"user_id" gorm:"unique;type:varchar(255)"`
	User      *User          `json:"user" gorm:"foreignKey:UserID"`
	Specialty string         `json:"specialty"`

	Tickets []Ticket `gorm:"many2many:assignedTickets;"`
}

// BeforeCreate These functions are called before creating any Post
func (technician *Technician) BeforeCreate(_ *gorm.DB) (err error) {
	technician.ID = uuid.New().String()
	return
}
