package services_test

import (
	"log"
	"testing"

	"github.com/Jackk-Doe/basic-go-crud-api/database"

	"github.com/joho/godotenv"
)

/**
* Set up database, router and environment for Unit tests in TestMain()
**/
func TestMain(m *testing.M) {
	log.Println()
	log.Println("--> START Unit testing of Post Services functions...")
	log.Println()

	// Intialize .env loader
	loadEnvTest()

	// Initialise Sqlite database to be used for testing
	database.InitSQLite()

	m.Run()

	log.Println()
	log.Println("--> END Unit testing of Post Services functions...")
	log.Println()
}

// Since Unit uses .env.test, hence it needs different .env loader
func loadEnvTest() {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Error loading .env.test file")
	}
}
