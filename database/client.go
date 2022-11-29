package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"www.github.com/N4choWasTaken/Go-API-JWT/database/models"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}
