package services

import (
	"github.com/Xuduoteng/gomall/internal/pkg/logger"
	"github.com/Xuduoteng/gomall/internal/pkg/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Get_db() *gorm.DB {
	if mysql.DB == nil {
		panic(`😫: mysql.DB is nil`)
	}
	return mysql.DB
}

func Update_db() {
	logger := logger.LogrusLogger
	if mysql.DB == nil {
		panic(`😫: mysql.DB is nil`)
	}
	db = mysql.DB
	logger.Printf(`🍟: Successfully updated mysql.DB`)
}
