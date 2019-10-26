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
func (a *parameter) SetSimpleConfig(w http.ResponseWriter, r *http.Request) {

	mod := engine.SimpleConfigReq{}
	json.NewDecoder(r.Body).Decode(&mod)
	resp := a.Add(&mod)

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")

	DefaultRespon(w, hasil)
	return
}
