package model

type ImunifyCreateParams struct {
	Code  string `json:"code"`
	Note  string `json:"note"`
	Limit string `json:"limit"`
}

type ImunifyUpdateParams struct {
	Key     string `json:"key"`
	Note    string `json:"note"`
	Limit   string `json:"limit"`
	Servers string `json:"servers"`
}

type ImunifyResponse struct {
	Data struct {
		Key   string `json:"key"`
		Limit string `json:"limit"`
		Code  string `json:"code"`
	} `json:"data"`
}

type ImunifyDeleteParams struct {
	Key string `json:"key"`
}

type ImunifyDeleteResponse struct {
	baseModel
	Data []struct {
		Id  string `json:"id"`
		Key string `json:"key"`
	} `json:"data"`
}

type ImunifyListResponse struct {
	baseModel
	Data []struct {
		Key     string      `json:"key"`
		Type    int         `json:"type"`
		Code    string      `json:"code"`
		Limit   int         `json:"limit"`
		Servers int         `json:"servers"`
		Note    interface{} `json:"note"`
	} `json:"data"`
}

type ImunifyListServersParams struct {
	Key string `json:"key"`
}

type ImunifyListServersResponse struct {
	baseModel
	Data []struct {
		Id  string `json:"id"`
		Key string `json:"key"`
	} `json:"data"`
}
