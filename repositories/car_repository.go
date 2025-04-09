package repositories

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
	"rental.com/api/db"
	"rental.com/api/models"
)

func Save(car *models.Car) error {
	result := db.DB.Create(car)
	return result.Error
}

func Update(car *models.Car) error {
	result := db.DB.Save(car)
	return result.Error
}

func Delete(car *models.Car) error {
	result := db.DB.Delete(car)
	return result.Error
}

func GetAll(params models.CarQueryParams) ([]models.Car, int64, error) {
	var cars []models.Car
	var totalCount int64

	query := db.DB.Model(&models.Car{}).
		// Joins("LEFT JOIN users ON users.id = cars.user_id"). // in caz ca e nevoie de sortare dupa user
		Preload("User")

	query = ApplyCarFilters(query, params.Filters)
	query = ApplyCarSort(query, params.Sort)

	err := query.Count(&totalCount).Error
	if err != nil {
		return nil, 0, fmt.Errorf("Error fetching total count: %v", err)
	}

	err = query.Limit(params.Pagination.PageSize).Offset((params.Pagination.Page - 1) * params.Pagination.PageSize).Find(&cars).Error
	if err != nil {
		return nil, 0, fmt.Errorf("Error fetching cars: %v", err)
	}

	return cars, totalCount, nil
}

func FindByUserID(userID int64) ([]models.Car, error) {
	var cars []models.Car
	result := db.DB.Where("user_id = ?", userID).Find(&cars)
	return cars, result.Error
}

func ApplyCarFilters(query *gorm.DB, filters models.CarFilters) *gorm.DB {
	val := reflect.ValueOf(filters)
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		switch v := value.(type) {
		case string:
			if v != "" {
				query = query.Where(fmt.Sprintf("cars.%s = ?", field.Name), v)

			}
		case int64:
			if value != int64(0) {
				query = query.Where(fmt.Sprintf("cars.%s = ?", field.Name), value)
			}
		default:
			continue
		}
	}

	return query
}

func ApplyCarSort(query *gorm.DB, sort []string) *gorm.DB {
	for _, s := range sort {
		parts := strings.Split(s, ":")
		if len(parts) != 2 {
			continue // ignori valorile incorecte
		}

		field := parts[0]
		order := strings.ToLower(parts[1])

		if order == "asc" || order == "desc" {
			query = query.Order(fmt.Sprintf("cars.%s %s", field, order))
		}
	}
	return query
}
