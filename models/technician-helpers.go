package models

type TechnicianCreate struct {
	Specialty string `json:"specialty" validate:"required"`
	UserID    string `json:"userId" validate:"required"`
}

type TechnicianEdit struct {
	ID        string `json:"id"`
	Specialty string `json:"specialty" validate:"required"`
	UserID    string `json:"userId" validate:"required"`
}
