package controller

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App is model for structure Application
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}
