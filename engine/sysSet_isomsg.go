package engine

import (
	"M-GateDBConfig/model"
	"M-GateDBConfig/provider/tools"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type (

	// IsoMsgSettingConfig model for Settings Message ISO using SystemSetting KEY_02
	IsoMsgSettingConfig struct {
		ID                      string
		BitActive               *string
		PrimaryAccountNumber    *string
		MerchantType            *string
		PosEntryMode            *string
		AcquiringInstitutionID  *string
		TerminalID              *string
		CardAcceptorID          *string
		CardAcceptorName        *string
		ForwardingInstitutionID *string
	}
)

// GetIsoMsgSetting configuration in DB KEY_01
func (s *systemsettings) GetIsoMsgSetting() ([]byte, *SysSettDefResp) {
	keyreg := "KEY_04"
	// get from DB
	reg := model.NewRegistry(keyreg, "")
	value, err := s.repository.Select(reg)
	if err != nil {
		fmt.Printf("%+v", err)
		if strings.ContainsAny("record not found", err.Error()) {
			return nil, &SysSettDefResp{
				ID:    "0",
				Error: "Error No Data Recorded",
			}
		}
		return nil, &SysSettDefResp{
			ID:    "0",
			Error: "Error Select to Registry Table",
		}
	}

	// decode string
	dec, _ := base64.StdEncoding.DecodeString(value.Value)

	// decrypt
	decripted, err1 := tools.Decrypt(keyEncrDecr, dec)
	if err1 != nil {
		fmt.Printf("%+v", err)
		return nil, &SysSettDefResp{
			ID:    "0",
			Error: "Error Decrypt Value",
		}
	}

	return decripted, nil
}

func (s *systemsettings) SetIsoMsgSetting(e *IsoMsgSettingConfig) *SysSettDefResp {
	keyreg := "KEY_04"
	// check format
	check := s.checkTagIsoMsg(e)
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
		encripted, err := tools.Encrypt(keyEncrDecr, js)
		if err != nil {
			fmt.Printf("%+v", err)
			return &SysSettDefResp{
				ID:    string(e.ID),
				Error: "Error Encrypting",
			}
		}

		// encode to *string
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

	// DO UPDATE IF NO VALUE
	// encrypt
	encripted, err := tools.Encrypt(keyEncrDecr, js)
	if err != nil {
		fmt.Printf("%+v", err)
		return &SysSettDefResp{
			ID:    string(e.ID),
			Error: "Error Encrypting",
		}
	}

	// encode to *string
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

func (s *systemsettings) checkTagIsoMsg(x *IsoMsgSettingConfig) string {
	if x.AcquiringInstitutionID == nil {
		return "Tag AcquiringInstitutionID is null "
	}
	if x.BitActive == nil {
		return "Tag BitActive is null "
	}
	if x.CardAcceptorID == nil {
		return "Tag CardAcceptorID is null "
	}
	if x.CardAcceptorName == nil {
		return "Tag CardAcceptorName is null "
	}
	if x.ForwardingInstitutionID == nil {
		return "Tag ForwardingInstitutionID is null "
	}
	if x.MerchantType == nil {
		return "Tag MerchantType is null "
	}
	if x.PosEntryMode == nil {
		return "Tag PosEntryMode is null "
	}
	if x.PrimaryAccountNumber == nil {
		return "Tag PrimaryAccountNumber is null "
	}
	if x.TerminalID == nil {
		return "Tag TerminalID is null "
	}
	return ""
}
