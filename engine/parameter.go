package engine

type (
	// Parameter is the interface for interactor
	Parameter interface {
		Add(m *SimpleConfigReq) *SimpleConfigResp
		// List(c context.Context, r *ListGreetingsRequest) *ListGreetingsResponse
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
