package model

type BackupCreateParams struct {
	ServerId   string `json:"server_id"`
	Size       string `json:"size"`
	Datacenter string `json:"datacenter"`
}

type BackupResizeParams struct {
	BackupId string `json:"backup_id"`
	Delta    string `json:"delta"`
}

type BackupRemoveParams struct {
	BackupId string `json:"backup_id"`
}

type BackupResponse struct {
	baseModel
	Data []struct {
		ServerId   string `json:"serverId"`
		Size       int    `json:"size"`
		BackupId   int    `json:"backupId"`
		Datacenter string `json:"datacenter"`
	} `json:"data"`
}
