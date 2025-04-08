package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"rental.com/api/models"
	"rental.com/api/services"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	if err := services.RegisterUser(context.Background(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// Obținem token-ul JWT
	token, err := services.AuthenticateUser(context.Background(), credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	// Returnăm token-ul ca răspuns JSON
	c.JSON(http.StatusOK, gin.H{"token": token})
}
