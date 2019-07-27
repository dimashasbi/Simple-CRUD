package handler

import (
	"M-GateDBConfig/dbhandler"
	"M-GateDBConfig/fileconfiguration"
	"M-GateDBConfig/model"
)

// SetParam to input value to DB
func (a *AppHandler) SetParam(DB dbhandler.ISetDBValue) (bool, error) {

	simple, errfc := fileconfiguration.GetSimpleConfig()
	if errfc != nil {
		return false, errfc
	}

	data := &model.Parameters{
		Key:   simple.Key,
		Value: simple.Value,
	}

	err := DB.SetParameter(data)

	return true, err
}
