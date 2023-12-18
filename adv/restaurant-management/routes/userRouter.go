package routes

import (
	controller "dev.pradeep/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/users", controller.GetUsers())
	incommingRoutes.GET("/users/:user_id", controller.GetUsers())
	incommingRoutes.POST("/users/signup", controller.SignUp())
	incommingRoutes.POST("/users/login", controller.Login())
}
