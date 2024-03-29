package engine

import (
	"M-GateDBConfig/model"
	"fmt"
	"strings"
)

type (
	// RmvParameterReq for Request
	RmvParameterReq struct {
		ID    string
		Key   *string
		Value *string
	}
	// RmvParameterResp for Respon
	RmvParameterResp struct {
		ID    string
		Error string
	}
)

func (p *parameter) Remove(m *RmvParameterReq) *RmvParameterResp {
	check := p.checkTagRmv(m)
	if check != "" {
		return &RmvParameterResp{
			ID:    m.ID,
			Error: check,
		}
	}
	param := model.NewParameters(*m.Key, "")
	err := p.repository.Remove(param)
	if err != nil {
		fmt.Printf("%+v", err)
		if strings.ContainsAny("record not found", err.Error()) {
			return &RmvParameterResp{
				ID:    string(m.ID),
				Error: "Error No Data Recorded",
			}
		}
		return &RmvParameterResp{
			Error: "Error Delete in Parameter Table",
		}
	}
	return &RmvParameterResp{
		ID:    m.ID,
		Error: "",
	}
}

func (p *parameter) checkTagRmv(m *RmvParameterReq) string {

	if m.Key == nil || *m.Key == "" {
		return "Tag Key is missing or null "
	}

	return ""
}
