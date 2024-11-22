package models

import "gorm.io/gorm"

type Example struct {
	gorm.Model
	Name   string `json:"name"`                         // Name
	Status string `json:"status" gorm:"default:active"` // Status, active or inactive
}
