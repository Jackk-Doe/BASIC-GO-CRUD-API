package main

import (
	"jackk-doe/go-crud-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	// Load .env & Connect to DB, init() runs before main()
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// NOTE : Suggest debug with 'CompileDaemon -command="./go-crud-api"' command

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Test": "Hello Go!"})
	})

	r.Run()
}
