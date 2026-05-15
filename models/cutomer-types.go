package models

type CustomerQueryCondition struct {
	Active bool   `json:"activated"`
	Status Status `json:"status"`
}
