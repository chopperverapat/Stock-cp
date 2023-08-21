package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"server/model"
)

var dblog *gorm.DB

func GetDBlog() *gorm.DB {
	return dblog
}

func SetupDBlog() {

	database, err := gorm.Open(sqlite.Open("databaselog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connected databse")
	}

	database.AutoMigrate(&model.Log{})

	dblog = database

}
