package routes

import (
	"github.com/gin-gonic/gin"
	"rental.com/api/handlers"
	"rental.com/api/middlewares"
)

func RegisterCarRoutes(router *gin.Engine) {
	carGroup := router.Group("/cars")
	carGroup.Use(middlewares.Authenticate)

	carGroup.POST("/", handlers.CreateCar)
	carGroup.GET("/", handlers.GetAllCars)
	carGroup.DELETE("/:id", handlers.DeleteCar)

}
