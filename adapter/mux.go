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
		Router            *mux.Router
		muxParameter      *parameter
		muxSystemSettings *systemSettings
	}
)

// InitializeServer Application
func (a *Handler) InitializeServer(f engine.EnginesFactory) {
	// add Engine
	a.muxParameter = &parameter{f.NewParameter()}
	a.muxSystemSettings = &systemSettings{f.NewSystemSettings()}
	a.Router = mux.NewRouter()
	a.SetURL()
}

// SetURL for reloading
func (a *Handler) SetURL() {
	a.POST("/addparam", a.SetParam)
	a.POST("/allparam", a.GetAllParam)
	a.POST("/updparam", a.UpdateParam)
	a.POST("/rmvparam", a.RemoveParam)
	a.POST("/selparam", a.SelectParam)
	a.POST("/getfrontconf", a.GetFrontConfig)
	a.POST("/setfrontconf", a.SetFrontConfig)
	a.POST("/getbaseconf", a.GetBaseConfig)
	a.POST("/setbaseconf", a.SetBaseConfig)
	a.POST("/getbackconf", a.GetBackConfig)
	a.POST("/setbackconf", a.SetBackConfig)
	a.POST("/getisomsgconf", a.GetIsoMsgConfig)
	a.POST("/setisomsgconf", a.SetIsoMsgConfig)
}

// GET wraps the router for GET method
func (a *Handler) GET(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// POST wraps the router for POST method
func (a *Handler) POST(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// SetParam for  Parameter Mux
func (a *Handler) SetParam(w http.ResponseWriter, r *http.Request) {
	a.muxParameter.SetParam(w, r)
}

// GetAllParam for  Parameter Mux
func (a *Handler) GetAllParam(w http.ResponseWriter, r *http.Request) {
	a.muxParameter.GetAllParam(w, r)
}

// UpdateParam for  Parameter Mux
func (a *Handler) UpdateParam(w http.ResponseWriter, r *http.Request) {
	a.muxParameter.UpdateParam(w, r)
}

// RemoveParam for  Parameter Mux
func (a *Handler) RemoveParam(w http.ResponseWriter, r *http.Request) {
	a.muxParameter.RemoveParam(w, r)
}

// SelectParam for  Parameter Mux
func (a *Handler) SelectParam(w http.ResponseWriter, r *http.Request) {
	a.muxParameter.SelectParam(w, r)
}

// GetFrontConfig for  SystemSettings Mux
func (a *Handler) GetFrontConfig(w http.ResponseWriter, r *http.Request) {
	a.muxSystemSettings.GetFrontConfig(w, r)
}

// SetFrontConfig for  SystemSettings Mux
func (a *Handler) SetFrontConfig(w http.ResponseWriter, r *http.Request) {
	a.muxSystemSettings.SetFrontConfig(w, r)
}

// GetBaseConfig for  SystemSettings Mux
func (a *Handler) GetBaseConfig(w http.ResponseWriter, r *http.Request) {
	a.muxSystemSettings.GetBaseConfig(w, r)
}

// SetBaseConfig for  SystemSettings Mux
func (a *Handler) SetBaseConfig(w http.ResponseWriter, r *http.Request) {
	a.muxSystemSettings.SetBaseConfig(w, r)
}

// GetBackConfig for  SystemSettings Mux
func (a *Handler) GetBackConfig(w http.ResponseWriter, r *http.Request) {
	a.muxSystemSettings.GetBackConfig(w, r)
}

// SetBackConfig for  SystemSettings Mux
func (a *Handler) SetBackConfig(w http.ResponseWriter, r *http.Request) {
	a.muxSystemSettings.SetBackConfig(w, r)
}

// GetIsoMsgConfig for  SystemSettings Mux
func (a *Handler) GetIsoMsgConfig(w http.ResponseWriter, r *http.Request) {
	a.muxSystemSettings.GetIsoMsgConfig(w, r)
}

// SetIsoMsgConfig for  SystemSettings Mux
func (a *Handler) SetIsoMsgConfig(w http.ResponseWriter, r *http.Request) {
	a.muxSystemSettings.SetIsoMsgConfig(w, r)
}

// Run the app on it's router
func (a *Handler) Run(port string) {
	http.ListenAndServe(port, a.Router)
}
