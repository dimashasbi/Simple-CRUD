package engine

import (
	"M-GateDBConfig/model"
	"fmt"
	"strings"
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
		fmt.Printf("%+v", err)
		if strings.ContainsAny("record not found", err.Error()) {
			return &ListParameterResp{
				Error: "Error No Data Recorded",
			}
		}
		return &ListParameterResp{
			Error: "Error Find in Parameter Table",
		}
	}
	return &ListParameterResp{
		ParameterList: hasil,
		Error:         "",
	}
}
