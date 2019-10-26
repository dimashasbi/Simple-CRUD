package engine

import (
	"M-GateDBConfig/model"
	"context"
)

func (p *parameter) Add(c context.Context, m *SimpleConfigReq) *SimpleConfigResp {
	param := model.NewParameters(m.Key, m.Value)
	err := p.repository.Insert(c, param)
	return &SimpleConfigResp{
		ID:    string(param.ID),
		Error: err.Error(),
	}
}
