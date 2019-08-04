package dbhandler

import (
	"M-GateDBConfig/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// // DBHandler use for interface every object use this package
// type DBHandler interface {
// 	SetDBValue(db *gorm.DB, val string) error
// 	GetDBValue(db *gorm.DB, num int) (string, error)
// }

// ISetDBValue interface
type ISetDBValue interface {
	SetParameter(model.Parameters) error
	SetSystemSettings(model.SystemSettings) error
	SetCaAPIMessages(model.CaAPIMessages) error
	SetCaAPITransactions(model.CaAPITransactions) error
	SetIsoAPIMessages(mod model.IsoAPIMessages) error
}

// IGetDBValue interface
type IGetDBValue interface {
	GetParameter(mod *model.Parameters) (model.Parameters, error)
	GetSystemSettings(mod *model.SystemSettings) (model.SystemSettings, error)
	GetCaAPIMessages(mod *model.CaAPIMessages) (model.CaAPIMessages, error)
	GetCaAPITransactions(mod *model.CaAPIMessages) (model.CaAPITransactions, error)
	GetIsoAPIMessages(mod *model.CaAPIMessages) (model.IsoAPIMessages, error)
}

// DBHandler for db structure
type DBHandler struct {
	database *gorm.DB
}

// SetParameter is using for set configuration value on DB
func (db *DBHandler) SetParameter(data model.Parameters) error {
	// put model to Database
	result := db.database.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SetSystemSettings is using for set configuration value on DB
func (db *DBHandler) SetSystemSettings(data model.SystemSettings) error {
	// put model to Database
	result := db.database.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SetCaAPIMessages is using for set configuration value on DB
func (db *DBHandler) SetCaAPIMessages(data model.CaAPIMessages) error {
	// put model to Database
	result := db.database.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SetCaAPITransactions is using for set configuration value on DB
func (db *DBHandler) SetCaAPITransactions(data model.CaAPITransactions) error {
	// put model to Database
	result := db.database.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SetIsoAPIMessages is using for set configuration value on DB
func (db *DBHandler) SetIsoAPIMessages(data model.IsoAPIMessages) error {
	// put model to Database
	result := db.database.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetParameter is using for set configuration value on DB
func (db *DBHandler) GetParameter(data model.Parameters) (model.Parameters, error) {
	// put model to Database
	result := db.database.First(data, "code = ?", "L1212") // find product with code l1212
	if result.Error != nil {
		return _, result.Error
	}
	return
}
