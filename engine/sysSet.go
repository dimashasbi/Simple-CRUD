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
	keyEncrDecr = []byte(fmt.Sprintf("%16v", mypass))
	return &systemsettings{
		repository: f.NewSystemSettingRespository(),
	}
}

var (
	mypass      = "JAQUESTSKYWALKER" // should under 16
	keyEncrDecr []byte               // gonna be 16
)

// // IsoMessageConfig model for Settings Message ISO using SystemSetting KEY_02
// IsoMessageConfig struct {
// 	BitActive               string
// 	PrimaryAccountNumber    string
// 	MerchantType            string
// 	PosEntryMode            string
// 	AcquiringInstitutionID  string
// 	TerminalID              string
// 	CardAcceptorID          string
// 	CardAcceptorName        string
// 	ForwardingInstitutionID string
// }
