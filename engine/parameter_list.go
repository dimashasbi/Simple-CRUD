package engine

import (
	"M-GateDBConfig/model"
)

type (
	// ListParameterReq for Request
	ListParameterReq struct {
		Count int
	}
	// ListParameterResp for Respon
	ListParameterResp struct {
		ParameterList []*model.Parameters
		Error         string
	}
)

func (p *parameter) List() *ListParameterResp {
	hasil, err := p.repository.List()
	if err != nil {
		return &ListParameterResp{
			ParameterList: nil,
			Error:         err.Error(),
		}
	}
	return &ListParameterResp{
		ParameterList: hasil,
		Error:         "",
	}
}
