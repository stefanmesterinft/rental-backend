package main

import (
	"rental.com/api/db"
	"rental.com/api/routes"
)

func main() {
	dsn := "postgres://postgres:insightft@localhost:5433/car-rental"
	dbInstance := db.InitDB(dsn)

	r := routes.SetupRoutes(dbInstance)

	r.Run(":8080")
}
