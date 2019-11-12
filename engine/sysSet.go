package engine

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
	return &systemsettings{
		repository: f.NewSystemSettingRespository(),
	}
}

var (
	keyEncrDecr = "6368616e676520746869732070617373776f726420746f206120736563726574"
	// keyEncrDecr = "DIMASHASBI---JAQUEST SKYWALKER---JUGGERNAUT---PHANTOM ASSASIN---"

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
