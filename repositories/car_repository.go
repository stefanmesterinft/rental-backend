package repositories

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gorm.io/gorm"
	"rental.com/api/db"
	"rental.com/api/models"
)

func Save(car *models.Car) error {
	result := db.DB.Create(car)
	return result.Error
}

func GetAllCars(params models.CarQueryParams) ([]models.Car, int64, error) {
	var cars []models.Car
	var totalCount int64

	fmt.Print(params)

	query := db.DB.Model(&models.Car{}).Joins("LEFT JOIN users ON users.id = cars.user_id")
	query = ApplyCarFilters(query, params.Filters)
	query = ApplyCarSort(query, params.Sort)

	err := query.Count(&totalCount).Error
	if err != nil {
		return nil, 0, fmt.Errorf("Error fetching total count: %v", err)
	}

	fmt.Print(query)

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

	filtersJson, err := json.Marshal(filters)
	if err != nil {
		fmt.Printf("Error marshaling filters to JSON: %v", err)
	} else {
		// fmthează parametrii sub formă de JSON
		fmt.Printf("Applying filters: %s", filtersJson)
	}

	val := reflect.ValueOf(filters)
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		switch v := value.(type) {
		case string:
			if v != "" {
				query = query.Where(fmt.Sprintf("cars.%s = ?", field.Name), v)
				query = query.Debug()

			}
		case int, int64:
			if v != 0 {
				query = query.Where(fmt.Sprintf("cars.%s = ?", field.Name), v)
				query = query.Debug()
			}
		default:
			continue
		}
	}

	query = query.Debug()

	return query
}

func ApplyCarSort(query *gorm.DB, sort models.Sort) *gorm.DB {
	for field, order := range sort {
		if order == "asc" || order == "desc" {
			query = query.Order(fmt.Sprintf("cars.%s %s", field, order))
		}
	}
	return query
}
