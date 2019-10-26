package engine

type (
	// SimpleConfigReq for Request Input
	SimpleConfigReq struct {
		Key   string
		Value string
	}

	// FrontSettingConfigResp for Response Front Setting
	FrontSettingConfigResp struct {
		ClientID  string
		PassKey   string
		SecretKey string
	}

	// SimpleConfigResp for Response Front Simple Config
	SimpleConfigResp struct {
		ID    string
		Error string
	}

	// HeaderHTTPMessageResp for Response Header HTTP
	HeaderHTTPMessageResp struct {
		URLJavaMPAY string
		URLJavaSVA  string
	}

	// BackSettingsConfigResp for Response Back Setting
	BackSettingsConfigResp struct {
		URLJavaMPAY string
		URLJavaSVA  string
	}
)
