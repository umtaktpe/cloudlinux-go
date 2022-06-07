package model

type baseModel struct {
	Message string      `json:"message,omitempty"`
	Type    interface{} `json:"type,omitempty"`
	Success bool        `json:"success,omitempty"`
}

type StatusResponse struct {
	baseModel
	Data struct {
		DbRhnOnline      bool `json:"db_rhn_online"`
		DbClwebConnected bool `json:"db_clweb_connected"`
		DbClwebOnline    bool `json:"db_clweb_online"`
		IPServerReg      bool `json:"ip_server_reg"`
		Xmlrpc           bool `json:"xmlrpc"`
		RhnOverloaded    bool `json:"rhn_overloaded"`
		DbRhnConnected   bool `json:"db_rhn_connected"`
		RhnQueueRejects  int  `json:"rhn_queue_rejects"`
		Acronis          []struct {
			DcName string `json:"dc_name"`
			Status bool   `json:"status"`
		} `json:"acronis"`
	} `json:"data"`
}
