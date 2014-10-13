package main

import (
	"os"
	
	"github.com/rushtehrani/scotch/common/cfg"
	"github.com/rushtehrani/scotch/common/app"
	"github.com/rushtehrani/scotch/common/db"
	"github.com/rushtehrani/scotch/users"
)

func main() {
	
	// Configuration
	cfg.Set("db.driverName", "postgres")
	cfg.Set("db.dataSourceName", os.Getenv("DB_DATA_SOURCE_NAME"))
	
	// Database
	db.MustConnect(cfg.Get("db.driverName"), cfg.Get("db.dataSourceName"))

	// Application
	app := app.New()

	// Routes
	app.Get("/users/{id}", users.Get)
	app.Post("/users/", users.Create)
	app.Put("/users/{id}", users.Update)

	app.Serve(":8080")
}
