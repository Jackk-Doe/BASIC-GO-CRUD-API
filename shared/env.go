package shared

import (
	"log"
	"os"
)

func GetDBURL() string {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("Error : DB_URL in .env is not found")
	}
	return dbUrl
}
