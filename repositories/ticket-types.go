package repositories

import "time"

type TicketResponse struct {
	ID            string    `json:"id"`
	Title         string    `json:"name"`
	Description   string    `json:"description"`
	CustomerID    string    `json:"customerId"`
	CustomerName  string    `json:"customerName"`
	CustomerEmail string    `json:"customerEmail"`
	CustomerPhone string    `json:"customerPhone"`
	CustomerImage string    `json:"customerImage"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
