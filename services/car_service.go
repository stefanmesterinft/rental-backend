package services

import (
	"fmt"
	"log"

	"rental.com/api/models"
	"rental.com/api/repositories"
)

func CreateCar(car *models.Car) error {
	return repositories.Save(car)
}

func GetAllCars(params models.CarQueryParams) ([]models.Car, int64, error) {
	log.Println("Entering GetAllCars service method...")

	// Folosim repository-ul pentru a interoga baza de date
	cars, totalCount, err := repositories.GetAllCars(params)
	if err != nil {
		log.Println("Error fetching cars:", err)
		return nil, 0, fmt.Errorf("Error fetching cars: %v", err)
	}

	return cars, totalCount, nil
}
