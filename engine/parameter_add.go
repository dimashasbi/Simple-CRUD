package engine

import (
	"M-GateDBConfig/model"
)

type (
	// AddConfigReq for Request Input
	AddConfigReq struct {
		ID    string
		Key   string
		Value string
	}

	// AddConfigResp for Response Front Simple Config
	AddConfigResp struct {
		ID    string
		Error string
	}
)

func (p *parameter) Add(m *AddConfigReq) *AddConfigResp {
	param := model.NewParameters(m.Key, m.Value)
	err := p.repository.Insert(param)
	if err != nil {
		return &AddConfigResp{
			ID:    string(m.ID),
			Error: err.Error(),
		}
	}
	return &AddConfigResp{
		ID:    string(m.ID),
		Error: "",
	}
}
