package services_test

import (
	"jackk-doe/go-crud-api/database"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

/**
* Set up database, router and environment for Unit tests in TestMain()
**/
func TestMain(m *testing.M) {
	log.Println("This is TestMain() START")

	// Intialize .env loader
	loadEnvTest()

	// Initialise Sqlite database to be used for testing
	database.InitSQLite()

	m.Run()

	log.Println("This is TestMain() END")
}

// Since Unit uses .env.test, hence it needs different .env loader
func loadEnvTest() {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Error loading .env.test file")
	}
}
