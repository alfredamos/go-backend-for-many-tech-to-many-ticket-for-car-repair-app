package models

type TicketQueryCondition struct {
	Completed  bool   `json:"completed"`
	Status     Status `json:"status"`
	CustomerID string `json:"customer_id"`
	Tech       string `json:"tech"`
}
