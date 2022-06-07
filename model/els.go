package model

type ELSCreateParams struct {
	Code  string `json:"code"`
	Note  string `json:"note"`
	Limit string `json:"limit"`
}

type ELSUpdateParams struct {
	Key   string `json:"key"`
	Limit string `json:"limit"`
	Note  string `json:"note"`
}

type ELSRemoveParams struct {
	Key string `json:"key"`
}

type ELSRemoveResponse struct {
	baseModel
	Data struct {
		Key string `json:"key"`
	} `json:"data"`
}

type ELSListParams struct {
	ProductType string `json:"productType"`
}

type ELSListResponse struct {
	baseModel
	Data []struct {
		Key         string `json:"key"`
		ProductCode string `json:"productCode"`
		UsageLimit  string `json:"usageLimit"`
		Description string `json:"description"`
	} `json:"data"`
}

type ELSListServersParams struct {
	ProductType string `json:"productType"`
}

type ELSListServersResponse struct {
	baseModel
	Data []struct {
		Id       string `json:"id"`
		Key      string `json:"key"`
		Ip       string `json:"ip"`
		Hostname string `json:"hostname"`
	} `json:"data"`
}

type ELSRemoveServerParams struct {
	Id string `json:"id"`
}

type ELSRemoveServerResponse struct {
	baseModel
	Data struct {
		Id       string `json:"id"`
		Key      string `json:"key"`
		Ip       string `json:"ip"`
		Hostname string `json:"hostname"`
	} `json:"data"`
}

type ELSResponse struct {
	baseModel
	Data struct {
		Key         string `json:"key"`
		ProductCode string `json:"productCode"`
		UsageLimit  string `json:"usageLimit"`
		Description string `json:"description"`
	} `json:"data"`
}
