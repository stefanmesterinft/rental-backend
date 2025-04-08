package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rental.com/api/models"
	"rental.com/api/services"
)

func CreateCar(context *gin.Context) {
	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userIDInt, ok := userId.(int64)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid user ID format"})
		return
	}

	var car models.Car
	if err := context.ShouldBindJSON(&car); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	car.UserID = userIDInt

	err := services.CreateCar(&car)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create car"})
		return
	}

	context.JSON(http.StatusCreated, car)
}

func GetAllCars(context *gin.Context) {
	var params models.CarQueryParams

	if err := context.ShouldBindQuery(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Parametrii invalizi", "details": err.Error()})
		return
	}

	if params.Pagination.Page == 0 {
		params.Pagination.Page = 1
	}
	if params.Pagination.PageSize == 0 {
		params.Pagination.PageSize = 10
	}

	cars, total, err := services.GetAllCars(params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la preluarea mașinilor", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"cars":       cars,
		"totalCount": total,
		"page":       params.Pagination.Page,
		"pageSize":   params.Pagination.PageSize,
	})
}
