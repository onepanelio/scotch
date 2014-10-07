package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	
	"github.com/rushtehrani/scotch/lib/db"
)

func main() {
	router := mux.NewRouter()
	
	db.Connect()
	
	// Routes
	router.HandleFunc("/users/{name}", getUser).Methods("GET")

	n := negroni.New()

	n.UseHandler(router)

	n.Run(":8080")
}
