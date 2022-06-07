package client

import "github.com/umtaktpe/cloudlinux-go/model"

// Create ELS key.
func (c *client) ELSCreate(params *model.ELSCreateParams) (*model.ELSResponse, error) {
	response := &model.ELSResponse{}
	if err := c.request("GET", "/els/key/create.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Update ELS key properties.
func (c *client) ELSUpdate(params *model.ELSUpdateParams) (*model.ELSResponse, error) {
	response := &model.ELSResponse{}
	if err := c.request("GET", "/els/key/update.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Remove ELS registration key with all servers.
func (c *client) ELSRemove(params *model.ELSRemoveParams) (*model.ELSRemoveResponse, error) {
	response := &model.ELSRemoveResponse{}
	if err := c.request("GET", "/els/key/remove.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// List all ELS keys owned by customer.
func (c *client) ELSList(params *model.ELSListParams) (*model.ELSListResponse, error) {
	response := &model.ELSListResponse{}
	if err := c.request("GET", "/els/key/list.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// List all ELS servers under specific key
func (c *client) ELSListServers(params *model.ELSListServersParams) (*model.ELSListServersResponse, error) {
	response := &model.ELSListServersResponse{}
	if err := c.request("GET", "/els/srv/list.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Remove ELS server by its ID.
func (c *client) ELSRemoveServer(params *model.ELSRemoveServerParams) (*model.ELSRemoveServerResponse, error) {
	response := &model.ELSRemoveServerResponse{}
	if err := c.request("GET", "/els/key/remove.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
