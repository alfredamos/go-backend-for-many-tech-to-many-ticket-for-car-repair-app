package models

type TicketStatus string

const (
	Closed TicketStatus = "Closed"
	Open   TicketStatus = "Open"
)

type AssignedTicketCreate struct {
	TechnicianID string `validate:"required" json:"techId"`
	TicketID     string `validate:"required" json:"ticketId"`
	AssignBy     string `validate:"required" json:"assignBy"`
}

type AssignedTicketEdit struct {
	TechnicianID string       `validate:"required" json:"techId"`
	TicketID     string       `validate:"required" json:"ticketId"`
	AssignBy     string       `validate:"required" json:"assignBy"`
	Status       TicketStatus `json:"status"`
	Completed    bool         `json:"completed"`
}
