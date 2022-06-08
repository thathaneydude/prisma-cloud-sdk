package cspm

import (
	"github.com/sirupsen/logrus"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/client"
	"net/http"
	"net/url"
)

type CspmClient struct {
	baseUrl    string
	baseClient client.BaseClient
}

// Get allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CspmClient) Get(endpoint string, params url.Values) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest(c.baseUrl, http.MethodGet, endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

// Post allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CspmClient) Post(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest(c.baseUrl, http.MethodPost, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

// Put allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CspmClient) Put(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest(c.baseUrl, http.MethodPut, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

// Patch allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CspmClient) Patch(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest(c.baseUrl, http.MethodPatch, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

// Delete allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CspmClient) Delete(endpoint string, params url.Values) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest(c.baseUrl, http.MethodDelete, endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

func (c *CspmClient) getWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Get(endpoint, params)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CspmClient) postWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Post(endpoint, body)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, &response)
}

func (c *CspmClient) putWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Put(endpoint, body)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CspmClient) patchWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Patch(endpoint, body)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CspmClient) deleteWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Delete(endpoint, params)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CspmClient) cspmDoWithRetry(req http.Request, currentAttempt int) (*http.Response, error) {
	resp, err := c.baseClient.DoWithRetry(req, currentAttempt)
	sErr, _ := err.(*client.UnauthorizedError)
	if sErr != nil {
		logrus.Debugf("Auth token may have expired. Attempting to refresh token")
		_, err = c.ExtendAuthToken()
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

// SetHeader set a header in the base client to a value provided
func (c *CspmClient) SetHeader(headerName string, headerValue string) {
	c.baseClient.Headers.Set(headerName, headerValue)
}

// OverwriteBaseClient allows to point to a different BaseClient if needed.
func (c *CspmClient) OverwriteBaseClient(b *client.BaseClient) {
	c.baseClient = *b
}

func (c *CspmClient) getBaseUrl() string {
	return c.baseUrl
}
