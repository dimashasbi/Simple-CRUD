package engine

import (
	"M-GateDBConfig/model"
	"M-GateDBConfig/provider/tools"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type (
	// FrontSettingConfig for Response Front Setting put in System Settings
	FrontSettingConfig struct {
		ID               string
		TrxListeningPort *string
		ClientID         *string
		PassKey          *string
		SecretKey        *string
	}
)

var (
	keyreg      = "KEY_01"
	keyEncrDecr = "6368616e676520746869732070617373776f726420746f206120736563726574"
	// keyEncrDecr = "JAQUESTSKYWALKR"
)

// GetFrontSetting configuration in DB KEY_01
func (s *systemsettings) GetFrontSetting() ([]byte, *SysSettDefResp) {

	// get from DB
	reg := model.NewRegistry(keyreg, "")
	value, err := s.repository.Select(reg)
	if err != nil {
		fmt.Printf("%+v", err)
		return nil, &SysSettDefResp{
			ID:    "0",
			Error: "Error Select to Registry Table",
		}
	}

	// decode string
	dec, _ := base64.StdEncoding.DecodeString(value.Value)

	// decrypt
	key, _ := hex.DecodeString(keyEncrDecr)

	decripted, err1 := tools.Decrypt(key, dec)
	if err1 != nil {
		fmt.Printf("%+v", err)
		return nil, &SysSettDefResp{
			ID:    "0",
			Error: "Error Decrypt Value",
		}
	}

	return decripted, nil
}

func (s *systemsettings) SetFrontSetting(e *FrontSettingConfig) *SysSettDefResp {
	// check format
	check := s.checkTagFront(e)
	if check != "" {
		return &SysSettDefResp{
			ID:    e.ID,
			Error: check,
		}
	}
	js, _ := json.Marshal(e)

	// check KEY there or not, if there do update else do insert
	reg := model.NewRegistry(keyreg, "")
	sel, err := s.repository.Select(reg)
	if err != nil {
		fmt.Printf("%+v", err)
		return &SysSettDefResp{
			ID:    string(e.ID),
			Error: "Error Insert to Registry Table",
		}
	}

	result := &SysSettDefResp{}
	if sel.Value == "" {
		result = s.doInsert(e, js)
	} else {
		result = s.doUpdate(e, js)
	}
	return result
}

func (s *systemsettings) doUpdate(e *FrontSettingConfig, inp []byte) *SysSettDefResp {

	// encrypt
	key, _ := hex.DecodeString(keyEncrDecr)
	encripted, err := tools.Encrypt(key, inp)
	if err != nil {
		fmt.Printf("%+v", err)
		return &SysSettDefResp{
			ID:    string(e.ID),
			Error: "Error Encrypt Value",
		}
	}

	// encode to string
	enc := base64.StdEncoding.EncodeToString(encripted)

	// put on DB
	regis := model.NewRegistry(keyreg, enc)
	err = s.repository.Update(regis)
	if err != nil {
		fmt.Printf("%+v", err)
		return &SysSettDefResp{
			ID:    string(e.ID),
			Error: "Error insert to Registry Table",
		}
	}
	return &SysSettDefResp{
		ID:    string(e.ID),
		Error: "",
	}
}

func (s *systemsettings) doInsert(e *FrontSettingConfig, inp []byte) *SysSettDefResp {

	// encrypt
	key, _ := hex.DecodeString(keyEncrDecr)
	encripted, err := tools.Encrypt(key, inp)
	if err != nil {
		fmt.Printf("%+v", err)
		return &SysSettDefResp{
			ID:    string(e.ID),
			Error: "Error Encrypt Value",
		}
	}

	// encode to string
	enc := base64.StdEncoding.EncodeToString(encripted)

	// put on DB
	regis := model.NewRegistry(keyreg, enc)
	err = s.repository.Insert(regis)
	if err != nil {
		fmt.Printf("%+v", err)
		return &SysSettDefResp{
			ID:    string(e.ID),
			Error: "Error insert to Registry Table",
		}
	}
	return &SysSettDefResp{
		ID:    string(e.ID),
		Error: "",
	}
}

func (s *systemsettings) checkTagFront(x *FrontSettingConfig) string {
	if x.ClientID == nil {
		return "Tag ClientID is null "
	}
	if x.PassKey == nil {
		return "Tag PassKey is null "
	}
	if x.SecretKey == nil {
		return "Tag SecretKey is null "
	}
	if x.TrxListeningPort == nil {
		return "Tag TrxListeningPort is null "
	}
	return ""
}
