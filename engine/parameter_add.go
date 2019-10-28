package engine

import (
	"M-GateDBConfig/model"
	"fmt"
)

type (
	// AddConfigReq for Request Input
	AddConfigReq struct {
		ID    string
		Key   *string
		Value *string
	}

	// AddConfigResp for Response Front Simple Config
	AddConfigResp struct {
		ID    string
		Error string
	}
)

func (p *parameter) Add(m *AddConfigReq) *AddConfigResp {

	check := checkTagAdd(m)
	if check != "" {
		return &AddConfigResp{
			ID:    m.ID,
			Error: check,
		}
	}
	param := model.NewParameters(*m.Key, *m.Value)
	err := p.repository.Insert(param)
	if err != nil {
		fmt.Printf("%+v", err)
		return &AddConfigResp{
			ID:    string(m.ID),
			Error: "Error input to Parameter Table",
		}
	}
	return &AddConfigResp{
		ID:    string(m.ID),
		Error: "",
	}
}

func checkTagAdd(m *AddConfigReq) string {

	if m.Key == nil || *m.Key == "" {
		return "Tag Key is missing or null "
	}

	if m.Value == nil {
		return "Tag Value is missing or null "
	}
	return ""
}
