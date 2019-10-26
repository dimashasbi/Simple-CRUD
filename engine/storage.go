package engine

type (
	// ParameterRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store greetings
	ParameterRepository interface {
		Parameter
	}

	// StorageFactory is the interface that a storage
	// provider needs to implement so that the engine can
	// request repository instances as it needs them
	StorageFactory interface {
		// NewGreetingRepository returns a storage specific
		// GreetingRepository implementation
		NewParameterRepository() ParameterRepository
	}
)
