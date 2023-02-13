package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sanyogpatel-tecblic/BackEnd-UseCase/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderitems/:orderitem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderitems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderitems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderitems/:orderitem_id", controller.UpdateOrderItem())
}
