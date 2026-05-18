package models

type TicketCreate struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CustomerID  string `json:"customerId" validate:"required"`
}

type TicketEdit struct {
	ID          string `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CustomerID  string `json:"customerId" validate:"required"`
}
