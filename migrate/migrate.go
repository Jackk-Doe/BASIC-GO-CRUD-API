package main

import (
	"jackk-doe/go-crud-api/initializers"
	"jackk-doe/go-crud-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// NOTE : Run this file once, to create (initialize) Table in PostGres db
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
