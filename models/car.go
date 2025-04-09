package models

type Car struct {
	ID       int64  `gorm:"primaryKey"`
	Make     string `json:"make" binding:"required"`
	Model    string `json:"model" binding:"required"`
	Year     int64  `json:"year" binding:"required"`
	Location string `json:"location" binding:"required"`
	Price    int64  `json:"price" binding:"required"`
	UserID   int64  `gorm:"column:user_id"`
	User     User   `gorm:"foreignKey:UserID"`
}

type CarUpdate struct {
	Make     string `json:"make" binding:"required"`
	Model    string `json:"model" binding:"required"`
	Year     int64  `json:"year" binding:"required"`
	Location string `json:"location" binding:"required"`
	Price    int64  `json:"price" binding:"required"`
}

type CarFilters struct {
	Make     string `form:"make"`
	Model    string `form:"model"`
	Location string `form:"location"`
	Price    int64  `form:"price"`
	Year     int64  `form:"year"`
}

type CarQueryParams struct {
	Filters    CarFilters `form:"filters"`
	Pagination Pagination `form:"pagination"`
	Sort       []string   `form:"sort[]"`
}
