package cs

import (
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/cspm"
	"net/http"
	"net/url"
)

type CsClient struct {
	cspmClient *cspm.CspmClient
}

// There's probably a better way to do this... Problem for another day.

// Get allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CsClient) Get(endpoint string, params url.Values) (*http.Response, error) {
	return c.cspmClient.Get(endpoint, params)
}

// Post allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CsClient) Post(endpoint string, body []byte) (*http.Response, error) {
	return c.cspmClient.Post(endpoint, body)
}

// Put allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CsClient) Put(endpoint string, body []byte) (*http.Response, error) {
	return c.cspmClient.Put(endpoint, body)
}

// Patch allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CsClient) Patch(endpoint string, body []byte) (*http.Response, error) {
	return c.cspmClient.Patch(endpoint, body)
}

// Delete allows the client to skip automatic unmarshalling of the http response and return it as-is
func (c *CsClient) Delete(endpoint string, params url.Values) (*http.Response, error) {
	return c.cspmClient.Delete(endpoint, params)
}

func (c *CsClient) getWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Get(endpoint, params)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CsClient) postWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Post(endpoint, body)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CsClient) putWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Put(endpoint, body)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CsClient) patchWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Patch(endpoint, body)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}

func (c *CsClient) deleteWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Delete(endpoint, params)
	if err != nil {
		return err
	}
	return internal.UnmarshalResponse(resp, response)
}
