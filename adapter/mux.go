package adapter

import (
	"M-GateDBConfig/provider/postgres"

	"github.com/gorilla/mux"

	"net/http"
)

// AppHandler using for make a Route
type AppHandler struct {
	Router    *mux.Router
	DBHandler *postgres.DBHandler
}

// InitializeServer Application
func (a *AppHandler) InitializeServer(db *postgres.DBHandler) {
	a.Router = mux.NewRouter()
	a.DBHandler = db
	// a.IDB = a.DBHandler
	a.SetURL()

}

// SetURL for reloading
func (a *AppHandler) SetURL() {
	// a.GET("/params", a.GetParam)
	// a.POST("/params", a.SetSimpleConfig)
}

// GET wraps the router for GET method
func (a *AppHandler) GET(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *AppHandler) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Run the app on it's router
func (a *AppHandler) Run(port string) {
	http.ListenAndServe(port, a.Router)
}

// GetParam for router for GET method
// func (a *AppHandler) GetParam(w http.ResponseWriter, r *http.Request) {
// 	// val, err := engine.GetParam(a.Engine)
// 	// if err != nil {
// 	// 	out := ErrRespon{
// 	// 		Error: err.Error(),
// 	// 	}
// 		hasil, _ := json.Marshal(out)
// 		DefaultRespon(w, hasil)
// 		return
// 	}
// 	hasil, _ := json.Marshal(val)
// 	w.Header().Set("Content-Type", "application/json")

// 	DefaultRespon(w, hasil)
// 	return
// }

// SetSimpleConfig for router for POST method
// func (a *AppHandler) SetSimpleConfig(w http.ResponseWriter, r *http.Request) {
// 	M := json.Decoder

// 	val, err := engine.SetSimpleConfig(a.Engine)
// 	if err != nil {
// 		out := ErrRespon{
// 			Error: err.Error(),
// 		}
// 		hasil, _ := json.Marshal(out)
// 		DefaultRespon(w, hasil)
// 		return
// 	}
// 	hasil, _ := json.Marshal(val)
// 	w.Header().Set("Content-Type", "application/json")

// 	DefaultRespon(w, hasil)
// 	return
// }
