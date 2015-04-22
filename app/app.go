package app

import (
	"net/http"

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
	a.chain = a.chain.Append(constructors...)
}

func (a *App) Head(path string, f func(http.ResponseWriter, *http.Request)) {
	a.HandleFunc(path, f).Methods("Head")
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

func (a *App) Delete(path string, f func(http.ResponseWriter, *http.Request)) {
	a.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) Options(path string, f func(http.ResponseWriter, *http.Request)) {
	a.HandleFunc(path, f).Methods("OPTIONS")
}

func (a *App) Listen(addr string) error {
	http.Handle("/", a.chain.Then(a))

	return http.ListenAndServe(addr, nil)
}
