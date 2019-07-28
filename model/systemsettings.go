package model

// YOU SHOULD NOT MAKE AN IMPORT INTERNAL PACKAGE IN MODEL!!!!!!!

// DBConfigurationModel file for Connecting DB
type DBConfigurationModel struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

// BaseApplicationConfig model for Configuration Standard
type BaseApplicationConfig struct {
	LogFolder   string
	CaName      string
	ActivateSVA string
	PlnDirect   string
	Database    string
}

// FrontSettingsConfig model for Settings Front Message / Parameter Configuration (url / message / parameter settings)
type FrontSettingsConfig struct {
	HTTPListeningPort string
	HeaderMessage     HeaderHTTPMessage
}

// HeaderHTTPMessage model for HTTP Message configuration
type HeaderHTTPMessage struct {
	ClientID  string
	PassKey   string
	SecretKey string
}

// BackSettingsConfig model for Settings Back Message / Parameter Configuration (url / message / parameter settings)
type BackSettingsConfig struct {
	URLJavaMPAY string
	URLJavaSVA  string
}

// IsoMessageConfig model for Settings Message ISO
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

// SimpleConfig model for Settings simple Configuration
type SimpleConfig struct {
	Key   string
	Value string
}
