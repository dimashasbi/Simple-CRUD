package adapter

import (
	"M-GateDBConfig/engine"

	"github.com/gorilla/mux"

	"net/http"
)

// Handler using for make a Route
type (
	// Handler structure for Application Start Server
	Handler struct {
		Router       *mux.Router
		muxParameter *parameter
	}
)

// InitializeServer Application
func (a *Handler) InitializeServer(f engine.EngineFactory) {
	a.muxParameter = &parameter{f.NewParameter()}
	a.Router = mux.NewRouter()
	a.SetURL()
}

// SetURL for reloading
func (a *Handler) SetURL() {
	a.POST("/AddParam", a.SetParam)
}

// GET wraps the router for GET method
func (a *Handler) GET(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// POST wraps the router for POST method
func (a *Handler) POST(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// SetSimpleConfig for  Parameter Mux
func (a *Handler) SetParam(w http.ResponseWriter, r *http.Request) {
	a.muxParameter.SetParam(w, r)
}

// SetSimpleConfig for  Parameter Mux
func (a *Handler) GetAllParam(w http.ResponseWriter, r *http.Request) {
	a.muxParameter.GetAllParam(w, r)
}

// Run the app on it's router
func (a *Handler) Run(port string) {
	http.ListenAndServe(port, a.Router)
}
