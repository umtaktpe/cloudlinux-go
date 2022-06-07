package model

type KernelCareCreateParams struct {
	Limit string `json:"limit"`
	Note  string `json:"note"`
}

type KernelCareCreateResponse struct {
	baseModel
	Data string `json:"data"`
}

type KernelCareDeleteParams struct {
	Key string `json:"key"`
}

type KernelCareDeleteResponse struct {
	baseModel
	Data bool `json:"data"`
}

type KernelCareListResponse struct {
	baseModel
	Data []struct {
		Key     string `json:"key"`
		Enabled bool   `json:"enabled"`
		Created string `json:"created"`
		Limit   int    `json:"limit"`
		Note    string `json:"note"`
	} `json:"data"`
}

type KernelCareServersParams struct {
	Key string `json:"key"`
}

type KernelCareServersResponse struct {
	baseModel
	Data []struct {
		ServerId string `json:"server_id"`
		Ip       string `json:"ip"`
		Created  string `json:"created"`
	} `json:"data"`
}
