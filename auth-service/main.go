package main

import (
	"net/http"

	"randogo.com/controllers"
	"randogo.com/initializers"
	"randogo.com/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()

	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello-world-to-docker")
	})
	router.Run()
}
