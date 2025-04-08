package routes

import (
	"github.com/gin-gonic/gin"
	"rental.com/api/handlers"
)

func RegisterUserRoutes(router *gin.Engine) {
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
}
