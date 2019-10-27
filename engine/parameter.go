package engine

type (
	// Parameter is the interface for interactor
	Parameter interface {
		Add(m *AddConfigReq) *AddConfigResp
		List() *ListParameterResp
		// parameter.Select()
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
