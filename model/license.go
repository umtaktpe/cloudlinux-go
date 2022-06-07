package model

type LicenseAvailabilityParams struct {
	Ip string `json:"ip"`
}

type LicenseAvailabilityResponse struct {
	baseModel
	Data struct {
		Available []int `json:"available"`
		Owned     []int `json:"owned"`
	} `json:"data"`
}

type LicenseCheckParams struct {
	Ip string `json:"ip"`
}

type LicenseCheckResponse struct {
	baseModel
	Data []int `json:"data"`
}

type LicenseRegisterParams struct {
	Ip         string `json:"ip"`
	Type       string `json:"type"`
	ElsAllowed string `json:"els_allowed"`
}

type LicenseRegisterResponse struct {
	baseModel
	Data struct {
		Ip         string `json:"ip"`
		Type       int    `json:"type"`
		Registered bool   `json:"registered"`
		Created    string `json:"created"`
	} `json:"data"`
}

type LicenseConvertParams struct {
	Ip     string `json:"ip"`
	Type   string `json:"type"`
	TypeTo string `json:"type_to"`
}

type LicenseConvertResponse struct {
	baseModel
	Data bool `json:"data"`
}

type LicenseRemoveParams struct {
	Ip   string `json:"ip"`
	Type string `json:"type"`
}

type LicenseRemoveResponse struct {
	baseModel
	Data bool `json:"data"`
}

type LicenseListResponse struct {
	Data []struct {
		Ip         string `json:"ip"`
		Type       int    `json:"type"`
		Registered bool   `json:"registered"`
		Created    string `json:"created"`
	} `json:"data"`
}

type LicenseServerParams struct {
	Ip string `json:"ip"`
}

type LicenseServerResponse struct {
	baseModel
	Data []struct {
		ServerInfo  string `json:"server_info"`
		Ip          string `json:"ip"`
		Type        int    `json:"type"`
		Registered  bool   `json:"registered"`
		Created     string `json:"created"`
		LastCheckin string `json:"last_checkin"`
	} `json:"data"`
}

type LicenseUpdateParams struct {
	Ip         string `json:"ip"`
	ElsAllowed string `json:"els_allowed"`
	Type       string `json:"type"`
}

type LicenseUpdateResponse struct {
	baseModel
	Data bool `json:"data"`
}
