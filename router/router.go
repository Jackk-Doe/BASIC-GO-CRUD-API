package router

import (
	"jackk-doe/go-crud-api/controllers"
	"jackk-doe/go-crud-api/shared"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// To set up [router] routes & port
func setUpRouter() {
	port := shared.GetPORT()

	router = gin.Default()

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Test": "Router connection test successed!"})
	})

	/// Post related routes
	postRouter := router.Group("/post")
	postRouter.GET("", controllers.PostGetAll)
	postRouter.GET("/:id", controllers.PostGetOneById)
	postRouter.POST("", controllers.PostCreate)
	postRouter.PUT("/:id", controllers.PostUpdate)
	postRouter.DELETE("/:id", controllers.PostDelete)

	/// Set up Port
	router.Run(":" + port)
}

// NOTE : is used in unit test file
func GetRouter() *gin.Engine {
	if router == nil {
		log.Fatal("Error : Router instance is not instanciated yet")
	}
	return router
}

func Init() {
	setUpRouter()
}