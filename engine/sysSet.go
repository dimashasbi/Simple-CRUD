package engine

type (
	// SystemSettings is the interface for interactor
	SystemSettings interface {
		GetFrontSetting() ([]byte, *SysSettDefResp)
		SetFrontSetting(m *FrontSettingConfig) *SysSettDefResp
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

// // BackSettingsConfigReq for Response Back Setting put in System Settings
// BackSettingsConfig struct {
// 	ID          string
// 	URLJavaMPAY string
// 	URLJavaSVA  string
// }

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
// // BaseApplicationConfig model for Configuration Standard
// BaseApplicationConfig struct {
// 	ActivateSVA string
// 	PlnDirect   string
// }
