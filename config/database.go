package config

import (
	"github.com/darwin1224/go-test/models"
	"github.com/jinzhu/gorm"
)

// InitDB -- Init mysql client
func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:rahasia@/go_test")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Article{})
	return db
}
