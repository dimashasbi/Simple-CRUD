package respon

import (
	"net/http"
)

// IRespon for Interfaceing build Respon
type IRespon interface {
	DefaultRespon(http.ResponseWriter, StRespon)
}

// StRespon for Struct Respon
type StRespon struct {
	Success bool
	Err     error
}

// DefaultRespon use for no Respon Build
func DefaultRespon(w http.ResponseWriter, St StRespon) {
	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if St.Err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(St.Err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
