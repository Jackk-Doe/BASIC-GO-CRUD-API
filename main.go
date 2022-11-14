package main

import (
	"github.com/gin-gonic/gin"
)

// Load .env file in init(), which run before main()
func init() {
	// TODO : load .env
	// TODO : initialize Database
}

// NOTE : Suggest debug with 'CompileDaemon -command="./go-crud-api"' command

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Test": "Hello Go!"})
	})

	r.Run()
}
