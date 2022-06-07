package cspm

import (
	"github.com/sirupsen/logrus"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	client2 "github.com/thathaneydude/prisma-cloud-sdk/internal/client"
	"net/http"
	"net/url"
)

type CspmClient struct {
	baseUrl    string
	BaseClient client2.BaseClientImpl
}

func (c *CspmClient) GetWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Get(endpoint, params)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CspmClient) PostWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Post(endpoint, body)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, &response)
}

func (c *CspmClient) PutWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Put(endpoint, body)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CspmClient) PatchWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Patch(endpoint, body)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CspmClient) DeleteWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Delete(endpoint, params)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CspmClient) Get(endpoint string, params url.Values) (*http.Response, error) {
	req, err := c.BaseClient.BuildRequest(c.baseUrl, http.MethodGet, endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

func (c *CspmClient) Post(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.BaseClient.BuildRequest(c.baseUrl, http.MethodPost, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

func (c *CspmClient) Put(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.BaseClient.BuildRequest(c.baseUrl, http.MethodPut, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

func (c *CspmClient) Patch(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.BaseClient.BuildRequest(c.baseUrl, http.MethodPatch, endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

func (c *CspmClient) Delete(endpoint string, params url.Values) (*http.Response, error) {
	req, err := c.BaseClient.BuildRequest(c.baseUrl, http.MethodDelete, endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

func (c *CspmClient) cspmDoWithRetry(req http.Request, currentAttempt int) (*http.Response, error) {
	resp, err := c.BaseClient.DoWithRetry(req, currentAttempt)
	sErr, _ := err.(*client2.UnauthorizedError)
	if sErr != nil {
		logrus.Debugf("Auth token may have expired. Attempting to refresh token")
		_, err = c.ExtendAuthToken()
		if err != nil {
			return nil, err
		}
		resp, err = c.BaseClient.Do(req)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CspmClient) SetHeader(headerName string, headerValue string) {
	c.BaseClient.Headers.Set(headerName, headerValue)
}

func (c *CspmClient) getBaseUrl() string {
	return c.baseUrl
}
