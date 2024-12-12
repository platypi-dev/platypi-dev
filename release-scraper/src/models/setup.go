package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("../repos.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}

	// perform database migration
	err = database.AutoMigrate(&Repo{})
	if err != nil {
		return
	}
	DB = database
}
