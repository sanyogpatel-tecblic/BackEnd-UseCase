package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sanyogpatel-tecblic/BackEnd-UseCase/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.GET("/foods/:food_id", controller.UpdateFood())
}
