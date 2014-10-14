package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type App struct {
	*mux.Router
	chain alice.Chain
}

func New() *App {
	return &App{Router: mux.NewRouter(), chain: alice.New()}
}

func (a *App) Use(constructors ...alice.Constructor) {
	a.chain.Append(constructors...)
}

func (a *App) Listen(addr string) {
	http.Handle("/", a.chain.Then(a))

	l := log.New(os.Stdout, "[app] ", 0)
	l.Printf("listening on %s", addr)
	l.Fatal(http.ListenAndServe(addr, nil))
}

func (a *App) Get(path string, f func(http.ResponseWriter, *http.Request)) {
	a.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(http.ResponseWriter, *http.Request)) {
	a.HandleFunc(path, f).Methods("POST")
}

func (a *App) Put(path string, f func(http.ResponseWriter, *http.Request)) {
	a.HandleFunc(path, f).Methods("PUT")
}
