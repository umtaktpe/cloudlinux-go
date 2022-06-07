package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/umtaktpe/cloudlinux-go"
	"github.com/umtaktpe/cloudlinux-go/utils"
)

const baseURL = "https://cln.cloudlinux.com/api"

type client struct {
	baseURL    string
	config     *cloudlinux.Config
	httpClient *http.Client
}

func NewClient(config *cloudlinux.Config) *client {
	return &client{
		baseURL: baseURL,
		config:  config,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *client) request(method, endpoint string, params, response interface{}) error {
	token := utils.GenerateToken(c.config.Username, c.config.SecretKey)
	url := fmt.Sprintf("%s%s?token=%s", c.baseURL, endpoint, token)
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	if params != nil {
		query := request.URL.Query()
		reflectVal := reflect.ValueOf(params).Elem()
		for i := 0; i < reflectVal.NumField(); i++ {
			value := reflectVal.Field(i).Interface().(string)
			if value != "" {
				query.Add(reflectVal.Type().Field(i).Tag.Get("json"), value)
			}
		}
		request.URL.RawQuery = query.Encode()
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}
