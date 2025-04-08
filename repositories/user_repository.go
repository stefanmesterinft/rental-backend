package repositories

import (
	"context"

	"rental.com/api/db"
	"rental.com/api/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	return db.DB.WithContext(ctx).Create(user).Error
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := db.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
