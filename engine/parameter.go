package engine

type (
	// Parameter is the interface for interactor
	Parameter interface {
		Add(m *AddParamReq) *AddParamResp
		List() *ListParameterResp
		Update(m *UpdParamReq) *UpdParamResp
		Remove(m *RmvParameterReq) *RmvParameterResp
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
