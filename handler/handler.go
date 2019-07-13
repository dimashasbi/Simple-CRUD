package handler

import (
	"M-GateDBConfig/configuration"
	"M-GateDBConfig/dbhandler"
	"M-GateDBConfig/model"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// HTTPHandler using for make a Route
type HTTPHandler struct {
	Router    *mux.Router
	DBHandler *dbhandler.DBHandler
}

// Initialize Application
func (a *HTTPHandler) Initialize(db *dbhandler.DBHandler) {
	a.Router = mux.NewRouter()
	a.DBHandler = db
	a.SetURL(db)
}

// SetURL for reloading
func (a *HTTPHandler) SetURL(db *dbhandler.DBHandler) {
	a.GET("/reload/params", a.ReloadParam) // reload param
}

// GET wraps the router for GET method
func (a *HTTPHandler) GET(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// ReloadParam for router for GET method
func (a *HTTPHandler) ReloadParam(w http.ResponseWriter, r *http.Request) {
	a.SetParam()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	fmt.Println("Input Data is Done")
}

// SetParam to input value to DB
func (a *HTTPHandler) SetParam() {

	simple := configuration.GetParamValue()

	data := &model.Parameters{
		Key:   simple.Key,
		Value: simple.Value,
	}

	err := a.DBHandler.SetParameter(data)
	checkErr(err)

}

// Run the app on it's router
func (a *HTTPHandler) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error in handler")
		// panic(err.Error())
	}
}

// base64.StdEncoding.EncodeToString(byte)
// base64.StdEncoding.DecodeString(string)

// // Read DB value
// decodestring, err := base64.StdEncoding.DecodeString(ans)
// checkErr(err)

// err = decrypt(key, decodestring)
// checkErr(err)
