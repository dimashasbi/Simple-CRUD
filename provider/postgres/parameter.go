package postgres

import (
	"M-GateDBConfig/engine"
	"M-GateDBConfig/model"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type (
	parameterRepository struct {
		session *gorm.DB
	}
)

func newParameterRepository(db *gorm.DB) engine.ParameterRepository {
	return &parameterRepository{db}
}

func (p parameterRepository) Insert(m *model.Parameters) error {
	result := p.session.Create(&m)
	if result.Error != nil {
		errStr := "Error input to Parameter Table \n"
		fmt.Printf(errStr, result.Error)
		return errors.Wrapf(result.Error, errStr)
	}
	return nil
}

func (p parameterRepository) List() ([]*model.Parameters, error) {
	mod := []*model.Parameters{}
	err := p.session.Find(&mod)
	if err != nil {
		errStr := "Error Find Parameter Table \n"
		fmt.Printf(errStr, err.Error)
		return mod, errors.Wrapf(err.Error, errStr)
	}
	return mod, nil
}

// func (p parameterRepository) Select(m *model.Parameters) ([]*model.Parameters, error) {

// }

// func (p parameterRepository) Update(m *model.Parameters) error {

// }
