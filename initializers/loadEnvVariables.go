package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// Load .env file found in this project
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Error loading .env file")
	}
}
