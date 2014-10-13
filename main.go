package main

import (
	"github.com/rushtehrani/scotch/common/app"
	"github.com/rushtehrani/scotch/common/db"
	"github.com/rushtehrani/scotch/users"
)

func main() {

	db.MustConnect()

	app := app.New()

	// Routes
	app.Get("/users/{id}", users.Get)
	app.Post("/users/", users.Create)
	app.Put("/users/{id}", users.Update)

	app.Serve(":8080")
}
