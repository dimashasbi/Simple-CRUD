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

	// BackSettingConfig for Response Back Setting put in System Settings
	BackSettingConfig struct {
		ID          string
		URLJavaMPAY *string
		URLJavaSVA  *string
	}
)

// GetBackSetting configuration in DB KEY_01
func (s *systemsettings) GetBackSetting() ([]byte, *SysSettDefResp) {
	keyreg := "KEY_03"
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

func (s *systemsettings) SetBackSetting(e *BackSettingConfig) *SysSettDefResp {
	keyreg := "KEY_03"
	// check format
	check := s.checkTagBack(e)
	if check != "" {
		return &SysSettDefResp{
			ID:    e.ID,
			Error: check,
		}
	}
	js, _ := json.Marshal(e)

	// check KEY there or not, if there do update else do insert
	reg := model.NewRegistry(keyreg, "")
	sel, _ := s.repository.Select(reg)

	if sel.Value == "" {
		// DO INSERT IF NO VALUE
		// encrypt
		key, _ := hex.DecodeString(keyEncrDecr)
		encripted, err := tools.Encrypt(key, js)
		if err != nil {
			fmt.Printf("%+v", err)
			return &SysSettDefResp{
				ID:    string(e.ID),
				Error: "Error Encrypting",
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

	// DO INSERT IF NO VALUE
	// encrypt
	key, _ := hex.DecodeString(keyEncrDecr)
	encripted, err := tools.Encrypt(key, js)
	if err != nil {
		fmt.Printf("%+v", err)
		return &SysSettDefResp{
			ID:    string(e.ID),
			Error: "Error Encrypting",
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

func (s *systemsettings) checkTagBack(x *BackSettingConfig) string {
	if x.URLJavaMPAY == nil {
		return "Tag URLJavaMPAY is null "
	}
	if x.URLJavaSVA == nil {
		return "Tag URLJavaSVA is null "
	}

	return ""
}
