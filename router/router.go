package router

import (
	"log"
	"net/http"

	"github.com/Jackk-Doe/basic-go-crud-api/controllers"
	"github.com/Jackk-Doe/basic-go-crud-api/middlewares"
	"github.com/Jackk-Doe/basic-go-crud-api/shared"

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
	postRouter.POST("", middlewares.AuthViaJWT, controllers.PostCreate)
	postRouter.PUT("/:id", middlewares.AuthViaJWT, controllers.PostUpdate)
	postRouter.DELETE("/:id", middlewares.AuthViaJWT, controllers.PostDelete)

	/// User related routes
	userRouter := router.Group("/user")
	userRouter.POST("/sign-up", controllers.UserSignUp)
	userRouter.POST("/log-in", controllers.UserLogIn)
	userRouter.GET("/me", middlewares.AuthViaJWT, controllers.GetMe)

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
