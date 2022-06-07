package client

import "github.com/umtaktpe/cloudlinux-go/model"

// Return all CloudLinux backups owned by authorized user.
func (c *client) BackupList() (*model.BackupResponse, error) {
	response := &model.BackupResponse{}
	if err := c.request("GET", "/ab/list.json", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Creates CloudLinux Backup for Imunify360 server
func (c *client) BackupCreate(params *model.BackupCreateParams) (*model.BackupResponse, error) {
	response := &model.BackupResponse{}
	if err := c.request("GET", "/ab/im/create.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Resizes CloudLinux Backup for Imunify360 server
func (c *client) BackupResize(params *model.BackupResizeParams) (*model.BackupResponse, error) {
	response := &model.BackupResponse{}
	if err := c.request("GET", "/ab/resize.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Removes CloudLinux Backup for Imunify360 server
func (c *client) BackupRemove(params *model.BackupRemoveParams) (*model.BackupResponse, error) {
	response := &model.BackupResponse{}
	if err := c.request("GET", "/ab/remove.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
