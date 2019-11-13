package engine

import "fmt"

type (
	// SystemSettings is the interface for interactor
	SystemSettings interface {
		GetFrontSetting() ([]byte, *SysSettDefResp)
		SetFrontSetting(m *FrontSettingConfig) *SysSettDefResp
		GetBaseSetting() ([]byte, *SysSettDefResp)
		SetBaseSetting(m *BaseSettingConfig) *SysSettDefResp
		GetBackSetting() ([]byte, *SysSettDefResp)
		SetBackSetting(m *BackSettingConfig) *SysSettDefResp
		GetIsoMsgSetting() ([]byte, *SysSettDefResp)
		SetIsoMsgSetting(m *IsoMsgSettingConfig) *SysSettDefResp
	}

	systemsettings struct {
		repository SystemSettingsRepository
	}

	// SysSettDefResp for default Respon after success Respon
	SysSettDefResp struct {
		ID    string
		Error string
	}
)

func (f *engineFactory) NewSystemSettings() SystemSettings {
	keyEncrDecr = []byte(fmt.Sprintf("%32v", mypass))
	return &systemsettings{
		repository: f.NewSystemSettingRespository(),
	}
}

var (
	mypass      = "JAQUESTSKYWALKERSS" // max 32
	keyEncrDecr []byte                 // gonna be 32
)
