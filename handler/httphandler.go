package handler

import (
	"M-GateDBConfig/dbhandler"
	"M-GateDBConfig/mylog"
	"M-GateDBConfig/respon"
	"fmt"
	"github.com/gorilla/mux"

	"net/http"
)

// AppHandler using for make a Route
type AppHandler struct {
	Router    *mux.Router
	IDB       dbhandler.ISetDBValue
	IRespon   respon.IRespon
	ILog      mylog.ILog
	DBHandler *dbhandler.DBHandler
	StRespon  respon.StRespon
}

// InitializeServer Application
func (a *AppHandler) InitializeServer(db *dbhandler.DBHandler) {
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

// Run the app on it's router
func (a *AppHandler) Run(port string) {
	ports := fmt.Sprintf("Running in Port %v", port)
	mylog.MnDebug(ports)
	
	mylog.FatalError(http.ListenAndServe(ports, a.Router),)
}

// ReloadParam for router for GET method
func (a *AppHandler) ReloadParam(w http.ResponseWriter, r *http.Request) {
	ok, err := a.SetParam(a.DBHandler)
	a.StRespon = respon.StRespon{
		Success: ok,
		Err:     err,
	}
	respon.DefaultRespon(w, a.StRespon)
	fmt.Println("Input Data is Done")
}
