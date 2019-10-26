package engine

import (
	"M-GateDBConfig/engine/parameterEngine"
)

type (
	// Parameter is interface for interactor
	Parameter interface {
		parameterEngine.AddSimpleConfig()
	}

	parameter struct {
		repository ParameterRepository
	}
)

func (f *engineFactory) NewParameter() Parameter {
	return &parameter{
		repository: f.NewParameterRepository(),
	}
}
