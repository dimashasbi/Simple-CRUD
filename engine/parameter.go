package engine

import (
	"context"
)

type (
	// Parameter is the interface for interactor
	Parameter interface {
		Add(c context.Context, m *SimpleConfigReq) *SimpleConfigResp
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
