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
	mod := engine.AddParamReq{}
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

func (a *parameter) UpdateParam(w http.ResponseWriter, r *http.Request) {
	mod := engine.UpdParamReq{}
	json.NewDecoder(r.Body).Decode(&mod)

	resp := a.Update(&mod)

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	DefaultRespon(w, hasil)
}

func (a *parameter) RemoveParam(w http.ResponseWriter, r *http.Request) {
	mod := engine.RmvParameterReq{}
	json.NewDecoder(r.Body).Decode(&mod)

	resp := a.Remove(&mod)

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	DefaultRespon(w, hasil)
}

func (a *parameter) SelectParam(w http.ResponseWriter, r *http.Request) {
	mod := engine.SelParamReq{}
	json.NewDecoder(r.Body).Decode(&mod)

	resp := a.Select(&mod)

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	DefaultRespon(w, hasil)
}
