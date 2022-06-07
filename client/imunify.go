package client

import "github.com/umtaktpe/cloudlinux-go/model"

// Create Imunify key.
func (c *client) ImunifyCreate(params *model.ImunifyCreateParams) (*model.ImunifyResponse, error) {
	response := &model.ImunifyResponse{}
	if err := c.request("GET", "/im/key/create.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Update Imunify key properties.
func (c *client) ImunifyUpdate(params *model.ImunifyUpdateParams) (*model.ImunifyResponse, error) {
	response := &model.ImunifyResponse{}
	if err := c.request("GET", "/im/key/update.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Remove Imunify registration key with all servers.
func (c *client) ImunifyDelete(params *model.ImunifyDeleteParams) (*model.ImunifyDeleteResponse, error) {
	response := &model.ImunifyDeleteResponse{}
	if err := c.request("GET", "/im/key/remove.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// List all Imunify keys owned by customer.
func (c *client) ImunifyList() (*model.ImunifyListResponse, error) {
	response := &model.ImunifyListResponse{}
	if err := c.request("GET", "/im/key/list.json", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// List all Imunify servers under specific key
func (c *client) ImunifyListServers(params *model.ImunifyListServersParams) (*model.ImunifyListServersResponse, error) {
	response := &model.ImunifyListServersResponse{}
	if err := c.request("GET", "/im/srv/list.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
