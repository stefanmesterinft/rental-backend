package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rental.com/api/models"
)

var DB *gorm.DB

func InitDB(dsn string) *gorm.DB {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Verifică dacă DB este nil
	if DB == nil {
		log.Fatal("Database instance is nil")
	}

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	log.Println("Database connected successfully")

	DB.AutoMigrate(&models.Car{}, &models.User{})
	return DB
}
