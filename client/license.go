package client

import "github.com/umtaktpe/cloudlinux-go/model"

// Will return information about what kind of license types are available for registration
// and what types are used by current account.
func (c *client) LicenseAvailability(params *model.LicenseAvailabilityParams) (*model.LicenseAvailabilityResponse, error) {
	response := &model.LicenseAvailabilityResponse{}
	if err := c.request("GET", "/ipl/availability.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Check if IP license is registered by any customer.
func (c *client) LicenseCheck(params *model.LicenseCheckParams) (*model.LicenseCheckResponse, error) {
	response := &model.LicenseCheckResponse{}
	if err := c.request("GET", "/ipl/check.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Will register IP based license for authorized user.
func (c *client) LicenseRegister(params *model.LicenseRegisterParams) (*model.LicenseRegisterResponse, error) {
	response := &model.LicenseRegisterResponse{}
	if err := c.request("GET", "/ipl/register.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Change Imunify IP license type.
func (c *client) LicenseConvert(params *model.LicenseConvertParams) (*model.LicenseConvertResponse, error) {
	response := &model.LicenseConvertResponse{}
	if err := c.request("GET", "/ipl/convert.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Will remove IP based license from authorized user licenses.
func (c *client) LicenseRemove(params *model.LicenseRemoveParams) (*model.LicenseRemoveResponse, error) {
	response := &model.LicenseRemoveResponse{}
	if err := c.request("GET", "/ipl/remove.json", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Return all IP licenses owned by authorized user.
func (c *client) LicenseList() (*model.LicenseListResponse, error) {
	response := &model.LicenseListResponse{}
	if err := c.request("GET", "/ipl/list.json", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Return all IP licenses owned by authorized user filtered by ip.
func (c *client) LicenseServer(params *model.LicenseServerParams) (*model.LicenseServerResponse, error) {
	response := &model.LicenseServerResponse{}
	if err := c.request("GET", "/ipl/server.json", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Update IP based license from authorized user licenses.
func (c *client) LicenseUpdate(params *model.LicenseUpdateParams) (*model.LicenseUpdateResponse, error) {
	response := &model.LicenseUpdateResponse{}
	if err := c.request("GET", "/ipl/update.json", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}
