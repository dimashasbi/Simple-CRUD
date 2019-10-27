package adapter

import (
	"M-GateDBConfig/engine"
	"encoding/json"
	"net/http"
)

type (
	parameter struct {
		engine.Parameter
	}
)

// SetSimpleConfig for router for POST method
func (a *parameter) SetParam(w http.ResponseWriter, r *http.Request) {

	mod := engine.AddConfigReq{}
	json.NewDecoder(r.Body).Decode(&mod)
	resp := a.Add(&mod)

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")

	DefaultRespon(w, hasil)
	return
}

// GetAllParam for router for POST method
func (a *parameter) GetAllParam(w http.ResponseWriter, r *http.Request) {
	resp := a.List()

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")

	DefaultRespon(w, hasil)
	return
}
