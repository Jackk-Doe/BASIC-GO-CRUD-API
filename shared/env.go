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

func GetPORT() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "3000"
	}
	return port
}

func GetSQLiteFile() string {
	dbFile := os.Getenv("SQLITE_FILE")
	if dbFile == "" {
		log.Fatal("Error : SQLITE_FILE in .env is not found")
	}
	return dbFile
}
