package models

type Pagination struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type Sort map[string]string
