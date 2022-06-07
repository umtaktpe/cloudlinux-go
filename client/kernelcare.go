package client

import "github.com/umtaktpe/cloudlinux-go/model"

// Will generate new KC key for authorized user.
func (c *client) KernelCareCreate(params *model.KernelCareCreateParams) (*model.KernelCareCreateResponse, error) {
	response := &model.KernelCareCreateResponse{}
	if err := c.request("GET", "/kcare/key/create.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Will delete KC key owned by authorized user.
func (c *client) KernelCareDelete(params *model.KernelCareDeleteParams) (*model.KernelCareDeleteResponse, error) {
	response := &model.KernelCareDeleteResponse{}
	if err := c.request("GET", "/kcare/key/delete.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Return list of all KC keys registered by authorized user.
func (c *client) KernelCareList() (*model.KernelCareListResponse, error) {
	response := &model.KernelCareListResponse{}
	if err := c.request("GET", "/kcare/key/delete.json", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Return list of servers registered with key owned by authorized user.
func (c *client) KernelCareServers(params *model.KernelCareServersParams) (*model.KernelCareServersResponse, error) {
	response := &model.KernelCareServersResponse{}
	if err := c.request("GET", "/kcare/key/servers.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
