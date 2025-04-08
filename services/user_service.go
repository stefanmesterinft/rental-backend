package services

import (
	"context"
	"errors"

	"rental.com/api/models"
	"rental.com/api/repositories"
	"rental.com/api/utils"
)

// Înregistrare utilizator
func RegisterUser(ctx context.Context, user *models.User) error {
	// Hash parola folosind utilitarul existent
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return repositories.CreateUser(ctx, user)
}

// Autentificare utilizator și generare JWT
func AuthenticateUser(ctx context.Context, email, password string) (string, error) {
	user, err := repositories.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Verificare parolă cu utilitarul existent
	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generare JWT
	token, err := utils.GenerateToken(user.Email, int64(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
