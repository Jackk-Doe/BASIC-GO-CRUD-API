package main

import (
	"jackk-doe/go-crud-api/controllers"
	"jackk-doe/go-crud-api/initializers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	// Load .env & Connect to DB, init() runs before main()
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// NOTE : Suggest debug with 'CompileDaemon -command="./go-crud-api"' command

func main() {

	// Set Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Test": "Hello Go!"})
	})

	r.POST("/post", controllers.PostCreate)
	r.GET("/post", controllers.PostGetAll)
	r.GET("/post/:id", controllers.PostGetOneById)
	r.PUT("/post/:id", controllers.PostUpdate)

	r.Run(":" + port)
}
