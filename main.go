package main

import (
	"os"

	"github.com/rushtehrani/scotch/app"
	"github.com/rushtehrani/scotch/cfg"
	"github.com/rushtehrani/scotch/db"
	//"github.com/rushtehrani/scotch/middleware/auth"
	"github.com/rushtehrani/scotch/users"
)

func init() {
	// Configuration - cfg vars should only be set during application startup
	cfg.Set("db.driverName", "postgres")
	cfg.Set("db.dataSourceName", os.Getenv("DB_DATA_SOURCE_NAME"))

	cfg.Set("auth.privateKeyPath", os.Getenv("AUTH_PRIVATE_KEY_PATH"))
	cfg.Set("auth.publicKeyPath", os.Getenv("AUTH_PUBLIC_KEY_PATH"))
}

func main() {
	// Database
	db.MustConnect(cfg.Get("db.driverName"), cfg.Get("db.dataSourceName"))

	// Application
	app := app.New()

	//app.Use(auth.New)

	// Routes
	app.Get("/users/{id}", users.Get)
	app.Post("/users/", users.Create)
	app.Put("/users/{id}", users.Update)

	app.Listen(":" + os.Getenv("PORT"))
}
