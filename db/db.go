package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"server/model"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func SetupDB() {

	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connected databse")
	}

	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.Transaction{})

	db = database

}
