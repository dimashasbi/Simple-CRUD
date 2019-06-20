package dbsetup

import (
	"time"
)

//// Application Configuration

// SystemSettings is model for database configuration system
type SystemSettings struct {
	// ID    uint `gorm:"primary_key"`
	Key   string `gorm:"size:40"`
	Value string `gorm:"size:255"`
}

// Parameters for key need to showing in Database
type Parameters struct {
	ID    uint   `gorm:"primary_key"`
	Key   string `gorm:"size:60"`
	Value string `gorm:"size:255"`
}

// // RoutingSVA for configuration SVA List
// type RoutingSVA struct {
// }

//// Logging

// CaAPIMessages for Logging Message from CA to API
type CaAPIMessages struct {
}

// CaAPITransactions for Logging Payment Transaction of API
type CaAPITransactions struct {
	ID                     uint `gorm:"primary_key"`
	DateTime               time.Time
	RequestID              string
	AuthKey                string
	TransactionType        int
	ProductIndicator       string
	RequestData            string
	ResponseData           string
	CustomerNumber         string
	ResponseCode           string
	Status                 string
	Reversed               string
	TerminalID             string
	STAN                   string
	RRN                    string
	ResponseDescription    string
	TransmissionDateTime   time.Time
	TransactionDate        time.Time
	OriginalDataElement    string
	AcquiringInstitutionID string
	IsoMessages            string
}

// // CaSvaTransactions for Logging SVA Transaction of API
// type CaSvaTransactions struct {
// }

// // CaTransactions for Logging CaAPITransactions but full ISO messages
// type CaTransactions struct {
// 	ID                       uint `gorm:"primary_key"`
// 	DateTime                 time.Time
// 	RequestID                string
// 	MTI                      string
// 	PrimaryAccountNumber     string
// 	ProcessingCode           string
// 	TransactionType          int
// 	Amount                   int
// 	TransmissionDateTime     time.Time
// 	SystemTraceAuditNumber   string
// 	TransactionDateTime      time.Time
// 	LocalTransactionDate     string
// 	LocalTransactionTime     string
// 	MerchantType             string
// 	AcquiringInstitutionID   string
// 	ForwardInstitutionID     string
// 	RetrievalReferenceNumber string
// 	CardAcceptorTerminalID   string
// 	CardAcceptorID           string
// 	CurrencyCode             string
// 	ResponseCode             string
// 	FromAccountNumber        string
// 	TransactionIndicator     string
// 	ProductIndicator         string
// 	CustomerNumber           string
// 	OriginalDataElement      string
// 	SessionID                string
// 	Status                   string
// 	Reversed                 string
// 	AdditionalDataReq        string
// 	AdditionalDataResp       string
// }

// IsoAPIMessages for Logging Transaction between API and system behind
type IsoAPIMessages struct {
	ID                       uint `gorm:"primary_key"`
	DirectionCode            int
	DateTime                 time.Time
	MTI                      string
	PrimaryAccountNumber     string
	ProcessingCode           string
	TransactionAmount        int
	TransmissionDateTime     time.Time
	SystemTraceAuditNumber   string
	LocalTransactionDate     string
	LocalTransactionTime     string
	AcquiringInstitutionID   string
	ForwardInstitutionID     string
	RetrievalReferenceNumber string
	CardAcceptorTerminalID   string
	CardAcceptorID           string
	AdditionalData           string
	CurrencyCode             string
	ResponseCode             string
	FromAccountNumber        string
	TransactionIndicator     string
	OriginalDataElement      string
}
