package database

import (
	"jackk-doe/go-crud-api/models"
	"jackk-doe/go-crud-api/shared"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// [DB] global database instance
var db *gorm.DB

// Get the Database instance via this function
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Error : Database instance is not instanciated yet")
	}
	return db
}

func connectToDB() {
	dbUrl := shared.GetDBURL()
	var err error
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Error : Fail to connect to Database")
	}
}

func migrateDBModels() {
	db.AutoMigrate(&models.Post{})
}

func Init() {
	connectToDB()
	migrateDBModels()
}
