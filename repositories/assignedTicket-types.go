package repositories

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
	"time"
)

type AssignedTicketResponse struct {
	TechID          string              `json:"techId"`
	TicketID        string              `json:"ticketId"`
	Status          models.TicketStatus `json:"status"`
	Completed       bool                `json:"completed"`
	AssignBy        string              `json:"assignBy"`
	AssignAt        time.Time           `json:"assignAt"`
	TicketTitle     string              `json:"ticketTitle"`
	TicketNotes     string              `json:"ticketNotes"`
	CustomerName    string              `json:"customerName"`
	CustomerEmail   string              `json:"customerEmail"`
	CustomerAddress string              `json:"customerAddress"`
	CustomerPhone   string              `json:"customerPhone"`
	CustomerImage   string              `json:"customerImage"`
	TechName        string              `json:"techName"`
	TechEmail       string              `json:"techEmail"`
	TechPhone       string              `json:"techPhone"`
	TechSpecialty   string              `json:"techSpecialty"`
	TechImage       string              `json:"techImage"`
}
