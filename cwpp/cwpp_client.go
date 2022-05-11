package cwpp

import (
	"github.com/prisma-cloud-sdk/client"
	"github.com/prisma-cloud-sdk/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type CwppClient struct {
	BaseClient client.BaseClientImpl
	consoleUrl string
	apiVersion string
	username   string
	password   string
}

func (c *CwppClient) GetWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Get(endpoint, params)
	if err != nil {
		return err
	}
	return utils.UnmarshalResponse(resp, response)
}

func (c *CwppClient) PostWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Post(endpoint, body)
	if err != nil {
		return err
	}
	return utils.UnmarshalResponse(resp, response)
}

func (c *CwppClient) PutWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Put(endpoint, body)
	if err != nil {
		return err
	}
	return utils.UnmarshalResponse(resp, response)
}

func (c *CwppClient) PatchWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Patch(endpoint, body)
	if err != nil {
		return err
	}
	return utils.UnmarshalResponse(resp, response)
}

func (c *CwppClient) DeleteWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Delete(endpoint, params)
	if err != nil {
		return err
	}
	return utils.UnmarshalResponse(resp, response)
}

func (c *CwppClient) Get(endpoint string, params url.Values) (*http.Response, error) {
	req, err := c.BaseClient.BuildRequest(c.consoleUrl, http.MethodGet, endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req)
}

func (c *CwppClient) Post(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.BaseClient.BuildRequest(c.consoleUrl, http.MethodPost, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req)
}

func (c *CwppClient) Put(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.BaseClient.BuildRequest(c.consoleUrl, http.MethodPut, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req)
}

func (c *CwppClient) Patch(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.BaseClient.BuildRequest(c.consoleUrl, http.MethodPatch, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req)
}

func (c *CwppClient) Delete(endpoint string, params url.Values) (*http.Response, error) {
	req, err := c.BaseClient.BuildRequest(c.consoleUrl, http.MethodDelete, endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req)
}

func (c *CwppClient) cwppDoWithRetry(req http.Request) (*http.Response, error) {
	resp, err := c.BaseClient.DoWithRetry(req, 1)
	sErr, _ := err.(*client.UnauthorizedError)
	if sErr != nil {
		logrus.Debugf("Auth token may have expired. Attempting to refresh token")
		_, err = c.Authenticate(c.username, c.password)
		if err != nil {
			return nil, err
		}
		resp, err = c.BaseClient.Do(req)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}
