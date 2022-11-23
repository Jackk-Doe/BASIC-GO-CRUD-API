package main

import (
	"jackk-doe/go-crud-api/database"
	"jackk-doe/go-crud-api/initializers"
	"jackk-doe/go-crud-api/router"
)

func main() {

	initializers.LoadEnvVariables()

	database.Init()
	router.Init()
}
