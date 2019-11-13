package engine

import (
	"M-GateDBConfig/model"
	"fmt"
	"strings"
)

type (
	// AddParamReq for Request Input
	AddParamReq struct {
		ID    string
		Key   *string
		Value *string
	}

	// AddParamResp for Response Front Simple Config
	AddParamResp struct {
		ID    string
		Error string
	}
)

func (p *parameter) Add(m *AddParamReq) *AddParamResp {

	check := p.checkTagAdd(m)
	if check != "" {
		return &AddParamResp{
			ID:    m.ID,
			Error: check,
		}
	}
	param := model.NewParameters(*m.Key, *m.Value)
	err := p.repository.Insert(param)
	if err != nil {
		fmt.Printf("%+v", err)
		if strings.ContainsAny("record not found", err.Error()) {
			return &AddParamResp{
				ID:    string(m.ID),
				Error: "Error No Data Recorded",
			}
		}
		return &AddParamResp{
			ID:    string(m.ID),
			Error: "Error input to Parameter Table",
		}
	}
	return &AddParamResp{
		ID:    string(m.ID),
		Error: "",
	}
}

func (p *parameter) checkTagAdd(m *AddParamReq) string {

	if m.Key == nil || *m.Key == "" {
		return "Tag Key is missing or null "
	}

	if m.Value == nil {
		return "Tag Value is missing or null "
	}
	return ""
}
