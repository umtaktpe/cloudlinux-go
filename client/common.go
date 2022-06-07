package client

import "github.com/umtaktpe/cloudlinux-go/model"

// Return system information about several Cloudlinux services.
func (c *client) Status() (*model.StatusResponse, error) {
	response := &model.StatusResponse{}
	if err := c.request("GET", "/status.json", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}
