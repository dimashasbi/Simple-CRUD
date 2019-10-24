package model

// DBConfigurationModel file for Connecting DB
type DBConfigurationModel struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

// // BaseApplicationConfig model for Configuration Standard
// type BaseApplicationConfig struct {
// 	LogFolder   string
// 	CaName      string
// 	ActivateSVA string
// 	PlnDirect   string
// }

// FrontSettingsConfig model for Settings Front Message / Parameter Configuration (url / message / parameter settings)
type FrontSettingsConfig struct {
	HTTPListeningPort string
	HeaderMessage     HeaderHTTPMessage
}

// HeaderHTTPMessage model for HTTP Message configuration using Table Parameter  using SystemSetting KEY_00
type HeaderHTTPMessage struct {
	ClientID  string
	PassKey   string
	SecretKey string
}

// BackSettingsConfig model for Settings Back Message / Parameter Configuration (url / message / parameter settings)  using SystemSetting KEY_01
type BackSettingsConfig struct {
	URLJavaMPAY string
	URLJavaSVA  string
}

// IsoMessageConfig model for Settings Message ISO using SystemSetting KEY_02
type IsoMessageConfig struct {
	BitActive               string
	PrimaryAccountNumber    string
	MerchantType            string
	PosEntryMode            string
	AcquiringInstitutionID  string
	TerminalID              string
	CardAcceptorID          string
	CardAcceptorName        string
	ForwardingInstitutionID string
}

// SimpleConfig model for Settings simple Configuration using Table Parameter
type SimpleConfig struct {
	Key   string
	Value string
}
