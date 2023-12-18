package routes

import (
	controller "dev.pradeep/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orderItems", controller.GetOrderItems())
	incommingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incommingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incommingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())

	incommingRoutes.GET("/orderItems/:order_id", controller.GetOrderItemsByOrder())
}
