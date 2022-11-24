package database

import (
	"log"

	"github.com/Jackk-Doe/basic-go-crud-api/models"
	"github.com/Jackk-Doe/basic-go-crud-api/shared"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
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

/**
* The belows are for Unit Tests related
*
* Note : for unit testing, create & use Sqlite as database, instead PostgreSQL
**/

func createSQLiteDB() {
	dbFile := shared.GetSQLiteFile()
	var err error
	db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatal("Error : Fail to open SQLite database")
	}
}

func InitSQLite() {
	createSQLiteDB()
	migrateDBModels()
}
