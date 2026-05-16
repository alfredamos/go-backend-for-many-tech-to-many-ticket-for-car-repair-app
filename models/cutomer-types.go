package models

type CustomerQueryCondition struct {
	Active bool   `json:"activated"`
	Status Status `json:"status"`
}

type CustomerCreate struct {
	Address string `json:"address" validate:"required"`
	Notes   string `json:"notes" validate:"required"`
	Active  bool   `json:"active" default:"true"`
	Status  Status
	UserID  string
}

type CustomerEdit struct {
	ID      string `json:"id"`
	Address string `json:"address" validate:"required"`
	Notes   string `json:"notes" validate:"required"`
	Active  bool   `json:"active" default:"true"`
	Status  Status
	UserID  string
}
