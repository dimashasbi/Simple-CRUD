package model

// model for Base Application Configuration
type BaseApplicationConfig struct {
	LogFolder   string
	CaName      string
	ActivateSVA string
	PlnDirect   string
	Database    string
}

// model for Settings Front Message / Parameter Configuration (url / message / parameter settings)
type FrontSettingsConfig struct {
	HTTPListeningPort string
	HeaderMessage     HeaderHTTPMessage
}

type HeaderHTTPMessage struct {
	ClientID  string
	PassKey   string
	SecretKey string
}

// model for Settings Back Message / Parameter Configuration (url / message / parameter settings)
type BackSettingsConfig struct {
	URLJavaMPAY string
	URLJavaSVA  string
}

// model for Settings Message ISO
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

// model for Settings simple Configuration
type SimpleConfig struct {
	Key   string
	Value string
}
