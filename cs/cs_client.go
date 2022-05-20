package cs

import (
	"github.com/thathaneydude/prisma-cloud-sdk/cspm"
	"net/http"
	"net/url"
)

type CsClient struct {
	cspmClient *cspm.CspmClient
}

// There's probably a better way to do this...

func (c *CsClient) GetWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	return c.cspmClient.GetWithResponseInterface(endpoint, params, response)
}

func (c *CsClient) PostWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	return c.cspmClient.PostWithResponseInterface(endpoint, body, response)
}

func (c *CsClient) PutWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	return c.cspmClient.PutWithResponseInterface(endpoint, body, response)
}

func (c *CsClient) PatchWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	return c.cspmClient.PatchWithResponseInterface(endpoint, body, response)
}

func (c *CsClient) DeleteWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	return c.cspmClient.DeleteWithResponseInterface(endpoint, params, response)
}

func (c *CsClient) Get(endpoint string, params url.Values) (*http.Response, error) {
	return c.cspmClient.Get(endpoint, params)
}

func (c *CsClient) Post(endpoint string, body []byte) (*http.Response, error) {
	return c.cspmClient.Post(endpoint, body)
}

func (c *CsClient) Put(endpoint string, body []byte) (*http.Response, error) {
	return c.cspmClient.Put(endpoint, body)
}

func (c *CsClient) Patch(endpoint string, body []byte) (*http.Response, error) {
	return c.cspmClient.Patch(endpoint, body)
}

func (c *CsClient) Delete(endpoint string, params url.Values) (*http.Response, error) {
	return c.cspmClient.Delete(endpoint, params)
}
