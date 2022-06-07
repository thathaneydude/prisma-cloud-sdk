package cwpp

import (
	"github.com/sirupsen/logrus"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/client"
	"net/http"
	"net/url"
)

type CwppClient struct {
	baseClient client.BaseClient
	consoleUrl string
	apiVersion string
	username   string
	password   string
}

// Get allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CwppClient) Get(endpoint string, params url.Values) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest(c.consoleUrl, http.MethodGet, endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req)
}

// Post allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CwppClient) Post(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest(c.consoleUrl, http.MethodPost, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req)
}

// Put allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CwppClient) Put(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest(c.consoleUrl, http.MethodPut, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req)
}

// Patch allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CwppClient) Patch(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest(c.consoleUrl, http.MethodPatch, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req)
}

// Delete allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CwppClient) Delete(endpoint string, params url.Values) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest(c.consoleUrl, http.MethodDelete, endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req)
}

func (c *CwppClient) getWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Get(endpoint, params)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CwppClient) postWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Post(endpoint, body)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CwppClient) putWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Put(endpoint, body)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CwppClient) patchWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Patch(endpoint, body)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CwppClient) deleteWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Delete(endpoint, params)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CwppClient) OverwriteBaseClient(b *client.BaseClient) {
	c.baseClient = *b
}

func (c *CwppClient) cwppDoWithRetry(req http.Request) (*http.Response, error) {
	resp, err := c.baseClient.DoWithRetry(req, 1)
	sErr, _ := err.(*client.UnauthorizedError)
	if sErr != nil {
		logrus.Debugf("Auth token may have expired. Attempting to refresh token")
		_, err = c.Authenticate(c.username, c.password)
		if err != nil {
			return nil, err
		}
		resp, err = c.baseClient.Do(req)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}
