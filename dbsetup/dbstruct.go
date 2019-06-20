package dbsetup

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

//// Application Configuration

// SystemSettings is model for database configuration system
type SystemSettings struct {
	// ID    uint `gorm:"primary_key"`
	Key   string `gorm:"size:40"`
	Value string `gorm:"size:2000"`
}

// Parameters for key need to showing in Database
type Parameters struct {
	ID    int    `gorm:"primary_key"`
	Key   string `gorm:"size:60"`
	Value string `gorm:"size:300"`
}

// // RoutingSVA for configuration SVA List
// type RoutingSVA struct {
// }

//// Logging

// CaAPIMessages for Logging Message from CA to API (1 line 1 direction)
type CaAPIMessages struct {
	ID           int `gorm:"primary_key"`
	DateTime     time.Time
	RequestID    string
	Direction    string
	RequestType  string
	RawMessage   postgres.Jsonb
	ResponseCode string `gorm:"size:2"`
	RRN          string
	IDCustomer   string
}

// CaAPITransactions for Logging Payment Transaction of API (1 line 2 direction)
type CaAPITransactions struct {
	ID                     int       `gorm:"primary_key;index:ID;AUTO_INCREMENT"`
	DateTime               time.Time `gorm:"not null"`
	RequestID              string    `gorm:"index:RequestID;not null"`
	AuthKey                string    `gorm:"index:AuthKey"`
	TransactionType        string    `gorm:"not null"`
	ProductIndicator       string    `gorm:"size:5"`
	RequestData            string
	ResponseData           string
	CustomerNumber         string `gorm:"index:IDCustomer;size:50"`
	ResponseCode           string `gorm:"not null;size:2"`
	Status                 string `gorm:"not null"`
	Reversed               string `gorm:"not null"`
	TerminalID             string `gorm:"index:TerminalID;size:16;"`
	STAN                   string `gorm:"index:STAN;size:6"`
	RRN                    string `gorm:"index:RRN;size:12"`
	ResponseDescription    string
	TransmissionDateTime   time.Time
	TransactionDateTime    time.Time
	OriginalDataElement    string `gorm:"size:42"`
	AcquiringInstitutionID string `gorm:"size:11"`
	IsoMessages            string `gorm:"not null"`
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

// IsoAPIMessages for Logging Transaction between API and system behind (1 line 1 direction)
type IsoAPIMessages struct {
	ID                       int       `gorm:"primary_key;index:ID;AUTO_INCREMENT"`
	DateTime                 time.Time `gorm:"not null"`
	IsoMessages              string    `gorm:"not null"`
	RequestID                string    `gorm:"index:RequestID;not null;"`
	MessageDestination       string    `gorm:"not null"`
	MTI                      string    `gorm:"not null;size:4"`
	PrimaryAccountNumber     string    `gorm:"not null;size:19"`
	ProcessingCode           string    `gorm:"not null;size:6"`
	TransactionAmount        int       `gorm:"not null;size:12"`
	TransmissionDateTime     time.Time `gorm:"not null"`
	SystemTraceAuditNumber   string    `gorm:"not null;size:6;index:STAN"`
	LocalTransactionDate     string    `gorm:"not null;size:6"`
	LocalTransactionTime     string    `gorm:"not null;size:4"`
	SettlementDate           string    `gorm:"not null;size:4"`
	MerchantType             string    `gorm:"not null;size:4"`
	PosEntryMode             string    `gorm:"not null;size:3"`
	AcquiringInstitutionID   string    `gorm:"not null;size:11"`
	ForwardInstitutionID     string    `gorm:"size:11"`
	Track2Data               string    `gorm:"size:37"`
	RetrievalReferenceNumber string    `gorm:"not null;size:12;index:RRN"`
	ResponseCode             string    `gorm:"not null;size:2"`
	CardAcceptorTerminalID   string    `gorm:"not null;size:16;index:TerminalID"`
	CardAcceptorID           string    `gorm:"not null;size:15;"`
	AdditionalData           string
	CurrencyCode             string `gorm:"not null;size:3"`
	FromAccountNumber        string `gorm:"size:28;"`
	TransactionIndicator     string
	OriginalDataElement      string `gorm:"size:42;"`
}
