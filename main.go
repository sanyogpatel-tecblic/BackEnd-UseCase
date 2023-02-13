package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/middleware"
	"github.com/sanyogpatel-tecblic/BackEnd-UseCase/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodcollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UseRoutes(router)
	router.Use(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.MenuRoutes(router)
	router.Run(":" + port)
}
