package routes

import (
	controller "dev.pradeep/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/menus", controller.GetMenus())
	incommingRoutes.GET("/menus/:menu_id", controller.GetMenu())
	incommingRoutes.POST("/menus", controller.CreateMenu())
	incommingRoutes.PATCH("/menus/:menu_id", controller.UpdateMenu())
}
