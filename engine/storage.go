package engine

import (
	"M-GateDBConfig/model"
)

type (
	// ParameterRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store greetings
	ParameterRepository interface {
		Insert(m *model.Parameters) error
		List() ([]*model.Parameters, error)
		// Select(m *model.Parameters) ([]*model.Parameters, error)
		Update(m *model.Parameters) error
	}

	// StorageFactory is the interface that a storage
	// provider needs to implement so that the engine can
	// request repository instances as it needs them
	StorageFactory interface {
		// NewParameterRepository returns a storage specific
		// GreetingRepository implementation
		NewParameterRepository() ParameterRepository
	}
)
