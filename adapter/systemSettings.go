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

func (h *systemSettings) GetBaseConfig(w http.ResponseWriter, r *http.Request) {
	resp, fail := h.GetBaseSetting()
	w.Header().Set("Content-Type", "application/json")
	if fail != nil {
		hasil, _ := json.Marshal(fail)
		DefaultRespon(w, hasil)
	} else {
		DefaultRespon(w, resp)
	}
}

func (h *systemSettings) SetBaseConfig(w http.ResponseWriter, r *http.Request) {
	mod := engine.BaseSettingConfig{}

	json.NewDecoder(r.Body).Decode(&mod)

	resp := h.SetBaseSetting(&mod)

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	DefaultRespon(w, hasil)
}

func (h *systemSettings) GetBackConfig(w http.ResponseWriter, r *http.Request) {
	resp, fail := h.GetBackSetting()
	w.Header().Set("Content-Type", "application/json")
	if fail != nil {
		hasil, _ := json.Marshal(fail)
		DefaultRespon(w, hasil)
	} else {
		DefaultRespon(w, resp)
	}
}

func (h *systemSettings) SetBackConfig(w http.ResponseWriter, r *http.Request) {
	mod := engine.BackSettingConfig{}

	json.NewDecoder(r.Body).Decode(&mod)

	resp := h.SetBackSetting(&mod)

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	DefaultRespon(w, hasil)
}

func (h *systemSettings) GetIsoMsgConfig(w http.ResponseWriter, r *http.Request) {
	resp, fail := h.GetIsoMsgSetting()
	w.Header().Set("Content-Type", "application/json")
	if fail != nil {
		hasil, _ := json.Marshal(fail)
		DefaultRespon(w, hasil)
	} else {
		DefaultRespon(w, resp)
	}
}

func (h *systemSettings) SetIsoMsgConfig(w http.ResponseWriter, r *http.Request) {
	mod := engine.IsoMsgSettingConfig{}

	json.NewDecoder(r.Body).Decode(&mod)

	resp := h.SetIsoMsgSetting(&mod)

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	DefaultRespon(w, hasil)
}
