package engine

type (
	// SimpleConfigReq for Request Input
	SimpleConfigReq struct {
		ID    string
		Key   string
		Value string
	}

	// FrontSettingConfigResp for Response Front Setting
	FrontSettingConfigResp struct {
		ID        string
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
		ID          string
		URLJavaMPAY string
		URLJavaSVA  string
	}

	// BackSettingsConfigResp for Response Back Setting
	BackSettingsConfigResp struct {
		ID          string
		URLJavaMPAY string
		URLJavaSVA  string
	}
)
