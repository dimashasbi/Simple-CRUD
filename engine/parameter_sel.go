package engine

import (
	"M-GateDBConfig/model"
	"fmt"
)

type (
	// SelParamReq for Request Input
	SelParamReq struct {
		ID    string
		Key   *string
		Value *string
	}

	// SelParamResp for Response Front Simple Config
	SelParamResp struct {
		ID    string
		Key   string
		Value string
		Error string
	}
)

func (p *parameter) Select(m *SelParamReq) *SelParamResp {

	check := p.checkTagSel(m)
	if check != "" {
		return &SelParamResp{
			ID:    m.ID,
			Error: check,
		}
	}
	param := model.NewParameters(*m.Key, "")
	result, err := p.repository.Select(param)
	if err != nil {
		fmt.Printf("%+v", err)
		return &SelParamResp{
			ID:    string(m.ID),
			Error: "Error Select to Parameter Table",
		}
	}
	return &SelParamResp{
		ID:    string(m.ID),
		Key:   *m.Key,
		Value: result.Value,
		Error: "",
	}
}

func (p *parameter) checkTagSel(m *SelParamReq) string {

	if m.Key == nil || *m.Key == "" {
		return "Tag Key is missing or null "
	}

	return ""
}
