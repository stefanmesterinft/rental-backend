package services

import (
	"fmt"

	"rental.com/api/models"
	"rental.com/api/repositories"
)

func CreateCar(car *models.Car) error {
	return repositories.Save(car)
}

func UpdateCar(existingCar *models.Car, input *models.CarUpdate) (*models.Car, error) {
	*existingCar = models.Car{
		ID:       existingCar.ID,
		UserID:   existingCar.UserID,
		Make:     input.Make,
		Model:    input.Model,
		Year:     input.Year,
		Location: input.Location,
		Price:    input.Price,
	}

	err := repositories.Update(existingCar)
	if err != nil {
		return nil, err
	}

	return existingCar, nil
}

func DeleteCar(car *models.Car) error {
	return repositories.Delete(car)
}

func GetAllCars(params models.CarQueryParams) ([]models.Car, int64, error) {
	cars, totalCount, err := repositories.GetAll(params)
	if err != nil {
		return nil, 0, fmt.Errorf("Error fetching cars: %v", err)
	}

	return cars, totalCount, nil
}
