package main

import (
	"go/jwt_auth/controllers"
	"go/jwt_auth/intializers"
	"go/jwt_auth/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectToMySQL()
	intializers.SyncDAtabase()
}

func main() {

	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth,controllers.Validate)

	r.Run()
}
