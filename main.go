package main

import (
	"rental.com/api/config"
	"rental.com/api/db"
	"rental.com/api/routes"
)

func main() {
	// Încarcă fișierul de configurare (.env)
	config.LoadConfig()

	// Obține string-ul de conexiune pentru DB
	dsn := config.GetDBConnectionString()

	// Inițializează conexiunea la baza de date
	dbInstance := db.InitDB(dsn)

	// Setup rutele
	r := routes.SetupRoutes(dbInstance)

	// Pornește serverul pe portul 8080
	r.Run(":8080")
}
