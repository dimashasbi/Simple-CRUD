package engine

import (
	"M-GateDBConfig/model"
)

func (p *parameter) Add(m *SimpleConfigReq) *SimpleConfigResp {
	param := model.NewParameters(m.Key, m.Value)
	err := p.repository.Insert(param)
	if err != nil {
		return &SimpleConfigResp{
			ID:    string(m.ID),
			Error: err.Error(),
		}
	}
	return &SimpleConfigResp{
		ID:    string(m.ID),
		Error: "",
	}
}
