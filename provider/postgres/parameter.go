package postgres

import (
	"M-GateDBConfig/engine"
	"M-GateDBConfig/model"

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
		return errors.Errorf("Error Input Parameter : %v", result.Error)
	}
	return nil
}

func (p parameterRepository) List() ([]*model.Parameters, error) {
	mod := []*model.Parameters{}
	result := p.session.Find(&mod)
	if result.Error != nil {
		return mod, errors.Errorf("Error Find Parameter Table : %v", result.Error)
	}
	return mod, nil
}

func (p parameterRepository) Update(m *model.Parameters) error {
	result := p.session.Model(&m).Where("KEY = ?", m.Key).Update("VALUE", m.Value)
	if result.Error != nil {
		return errors.Errorf("Error Update Parameter : %v", result.Error)
	}
	return nil
}

func (p parameterRepository) Remove(m *model.Parameters) error {
	result := p.session.Where("KEY = ?", m.Key).Delete(&m)
	if result.Error != nil {
		return errors.Errorf("Error Delete a Parameter : %v", result.Error)
	}
	return nil
}

func (p parameterRepository) Select(m *model.Parameters) (*model.Parameters, error) {
	result := p.session.Where("KEY = ?", m.Key).First(&m)
	if result.Error != nil {
		return m, errors.Errorf("Error Select a Parameter : %v", result.Error)
	}
	return m, nil
}
