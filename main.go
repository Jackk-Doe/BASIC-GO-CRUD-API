package main

import (
	"github.com/Jackk-Doe/basic-go-crud-api/database"
	"github.com/Jackk-Doe/basic-go-crud-api/initializers"
	"github.com/Jackk-Doe/basic-go-crud-api/router"
)

func main() {

	initializers.LoadEnvVariables()

	database.Init()
	router.Init()
}
