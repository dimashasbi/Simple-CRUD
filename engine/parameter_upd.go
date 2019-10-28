package engine

import (
	"M-GateDBConfig/model"
	"fmt"
)

type (
	// UpdParamReq struct Request
	UpdParamReq struct {
		ID    string
		Key   *string
		Value *string
	}

	// UpdParamResp struct Response
	UpdParamResp struct {
		ID    string
		Error string
	}
)

func (p *parameter) Update(m *UpdParamReq) *UpdParamResp {
	check := checkTagUpd(m)
	if check != "" {
		return &UpdParamResp{
			ID:    m.ID,
			Error: check,
		}
	}
	param := model.NewParameters(*m.Key, *m.Value)
	err := p.repository.Update(param)
	if err != nil {
		fmt.Printf("%+v", err)
		return &UpdParamResp{
			ID:    string(m.ID),
			Error: "Error Update to Parameter Table",
		}
	}
	return &UpdParamResp{
		ID:    string(m.ID),
		Error: "",
	}
}

func checkTagUpd(m *UpdParamReq) string {

	if m.Key == nil || *m.Key == "" {
		return "Tag Key is missing or null "
	}

	if m.Value == nil {
		return "Tag Value is missing or null "
	}
	return ""
}
