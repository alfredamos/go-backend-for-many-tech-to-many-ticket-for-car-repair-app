package repositories

import "go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

type CustomerQueryCondition struct {
	Active bool
	Status models.Status
	UserId string
}

type CustomerResponse struct {
	ID      string        `json:"id"`
	Address string        `json:"address"`
	Active  bool          `json:"active"`
	Notes   string        `json:"notes"`
	Status  models.Status `json:"status"`
	Name    string        `json:"name"`
	Email   string        `json:"email"`
	Phone   string        `json:"phone"`
	Image   string        `json:"image"`
	Gender  models.Gender `json:"gender"`
	Type    string        `json:"type"`
	UserId  string        `json:"userId"`
}
