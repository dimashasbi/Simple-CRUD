package engine

type (

	// FrontSettingConfigResp for Response Front Setting put in System Settings
	FrontSettingConfigResp struct {
		ID        string
		ClientID  string
		PassKey   string
		SecretKey string
	}

	// HeaderHTTPMessageReq for Response Header HTTP put in System Settings
	HeaderHTTPMessageReq struct {
		ID        string
		ClientID  string
		SecretKey string
	}

	// BackSettingsConfigReq for Response Back Setting put in System Settings
	BackSettingsConfigReq struct {
		ID          string
		URLJavaMPAY string
		URLJavaSVA  string
	}
)
