package controller

import (
	"MGateAllCA/controller/process"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (handler *App) Initialize() {
	handler.Router = mux.NewRouter()
	handler.setProcess()
}

func (a *App) setProcess() {
	a.GET("/api/billers", a.GetBiller)
}

// Get wraps the router for GET method
func (a *App) GET(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) POST(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) GetBiller(w http.ResponseWriter, r *http.Request) {
	process.GetBiller(w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
