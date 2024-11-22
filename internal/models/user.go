package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string     `json:"user_name"`     // Username
	Password    string     `json:"password"`      // Password
	Email       string     `json:"email"`         // Email
	LastLoginAt *time.Time `json:"last_login_at"` // Last login time
}
