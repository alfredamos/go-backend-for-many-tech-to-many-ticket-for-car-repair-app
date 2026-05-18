package repositories

import "go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

type TechnicianResponse struct {
	ID        string        `json:"id"`
	Specialty string        `json:"specialty"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Image     string        `json:"image"`
	Gender    models.Gender `json:"gender"`
	Type      string        `json:"type"`
	UserId    string        `json:"userId"`
}
