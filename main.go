package main

import (
	"rental.com/api/config"
	"rental.com/api/db"
	"rental.com/api/routes"
)

func main() {
	config.LoadConfig()
	dsn := config.GetDBConnectionString()

	dbInstance := db.InitDB(dsn)
	r := routes.SetupRoutes(dbInstance)

	r.Run(":8080")
}
