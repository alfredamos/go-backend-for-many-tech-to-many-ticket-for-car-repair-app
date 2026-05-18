package models

type CustomerQueryCondition struct {
	Active bool   `json:"activated"`
	Status Status `json:"status"`
}

type CustomerCreate struct {
	Address string `json:"address" validate:"required"`
	Notes   string `json:"notes" validate:"required"`
	Active  bool   `json:"active" gorm:"default:true"`
	Status  Status `json:"status" gorm:"type:varchar(20);default:'Invalid'"`
	UserID  string `json:"userId" validate:"required"`
}

type CustomerEdit struct {
	ID      string `json:"id"`
	Address string `json:"address" validate:"required"`
	Notes   string `json:"notes" validate:"required"`
	Active  bool   `json:"active" gorm:"default:true"`
	Status  Status `json:"status" gorm:"type:varchar(20);default:'Invalid'"`
	UserID  string `json:"userId" validate:"required"`
}

type CustomerResult struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	Notes   string `json:"notes"`
	Active  bool   `json:"active"`
	Status  Status `json:"status"`
	user    User
}
