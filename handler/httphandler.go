package handler

import (
	"M-GateDBConfig/dbhandler"
	"M-GateDBConfig/respon"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// AppHandler using for make a Route
type AppHandler struct {
	Router    *mux.Router
	DBHandler *dbhandler.DBHandler
	IDB       dbhandler.ISetDBValue
	IRespon   respon.IRespon
	StRespon  respon.StRespon
}

// Initialize Application
func (a *AppHandler) Initialize(db *dbhandler.DBHandler) {
	a.Router = mux.NewRouter()
	a.DBHandler = db
	a.SetURL(db)
}

// SetURL for reloading
func (a *AppHandler) SetURL(db *dbhandler.DBHandler) {
	a.GET("/reload/params", a.ReloadParam) // reload param
}

// GET wraps the router for GET method
func (a *AppHandler) GET(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// ReloadParam for router for GET method
func (a *AppHandler) ReloadParam(w http.ResponseWriter, r *http.Request) {
	ok, err := a.SetParam(a.DBHandler)
	resp := a.StRespon
	resp.Err = err
	resp.Success = ok
	respon.DefaultRespon(w, resp)
	fmt.Println("Input Data is Done")
}

// Run the app on it's router
func (a *AppHandler) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error in handler")
		panic(err.Error())
	}
}
