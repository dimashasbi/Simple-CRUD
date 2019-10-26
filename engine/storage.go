package engine

import (
	"M-GateDBConfig/model"
	"context"
)

type (
	// ParameterRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store greetings
	ParameterRepository interface {
		Insert(c context.Context, m *model.Parameters) error
		// List() error
		// Select() error
		// Update() error
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
