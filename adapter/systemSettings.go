package adapter

import (
	"M-GateDBConfig/engine"
	"encoding/json"
	"net/http"
)

type (
	systemSettings struct {
		engine.SystemSettings
	}
)

func (h *systemSettings) GetFrontConfig(w http.ResponseWriter, r *http.Request) {
	resp, fail := h.GetFrontSetting()

	w.Header().Set("Content-Type", "application/json")
	if fail != nil {
		hasil, _ := json.Marshal(fail)
		DefaultRespon(w, hasil)
	} else {
		DefaultRespon(w, resp)
	}
}

func (h *systemSettings) SetFrontConfig(w http.ResponseWriter, r *http.Request) {
	mod := engine.FrontSettingConfig{}

	json.NewDecoder(r.Body).Decode(&mod)

	resp := h.SetFrontSetting(&mod)

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	DefaultRespon(w, hasil)

}
