package handler

import (
	"M-GateDBConfig/dbhandler"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// AppHandler using for make a Route
type AppHandler struct {
	Router      *mux.Router
	DBHandler   *dbhandler.DBHandler
	DBInterface dbhandler.SetDBValueInterface
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
	err := a.SetParam(a.DBHandler)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("OK"))
	}
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

// base64.StdEncoding.EncodeToString(byte)
// base64.StdEncoding.DecodeString(string)

// // Read DB value
// decodestring, err := base64.StdEncoding.DecodeString(ans)
// checkErr(err)

// err = decrypt(key, decodestring)
// checkErr(err)
