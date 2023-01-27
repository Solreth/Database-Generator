package database

import (
	"Basic-Backend/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected successfully to the database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&models.User{})

	Database = DbInstance{Db: db}
}

type DbInstance struct {
	Db *gorm.DB
}
