package cspm

import (
	"PrismaCloud/pkg"
	bc "PrismaCloud/pkg/client"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

type CspmClient struct {
	baseClient bc.BaseClientImpl
}

func (c *CspmClient) GetWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Get(endpoint, params)
	if err != nil {
		return err
	}
	return c.unmarshalResponse(resp, response)
}

func (c *CspmClient) PostWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Post(endpoint, body)
	if err != nil {
		return err
	}
	return c.unmarshalResponse(resp, &response)
}

func (c *CspmClient) PutWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Put(endpoint, body)
	if err != nil {
		return err
	}
	return c.unmarshalResponse(resp, response)
}

func (c *CspmClient) PatchWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Patch(endpoint, body)
	if err != nil {
		return err
	}
	return c.unmarshalResponse(resp, response)
}

func (c *CspmClient) DeleteWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Delete(endpoint, params)
	if err != nil {
		return err
	}
	return c.unmarshalResponse(resp, response)
}

func (c *CspmClient) Get(endpoint string, params url.Values) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest("GET", endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

func (c *CspmClient) Post(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest("POST", endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

func (c *CspmClient) Put(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest("PUT", endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

func (c *CspmClient) Patch(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest("PATCH", endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

func (c *CspmClient) Delete(endpoint string, params url.Values) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest("DELETE", endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return c.cspmDoWithRetry(*req, 1)
}

func (c *CspmClient) cspmDoWithRetry(req http.Request, currentAttempt int) (*http.Response, error) {
	resp, err := c.baseClient.DoWithRetry(req, currentAttempt)
	sErr, _ := err.(*bc.UnauthorizedError)
	if sErr != nil {
		logrus.Debugf("Auth token may have expired. Attempting to refresh token")
		_, err = c.ExtendAuthToken()
		if err != nil {
			return nil, err
		}
		resp, err = c.baseClient.DoWithRetry(req, currentAttempt)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (c *CspmClient) unmarshalResponse(httpResponse *http.Response, response interface{}) error {
	if httpResponse == nil {
		return &pkg.GenericError{Msg: fmt.Sprintf("Error while reading response: No data found")}
	}

	defer httpResponse.Body.Close()
	tmp, readErr := io.ReadAll(httpResponse.Body)
	logrus.Debugf(string(tmp))
	if readErr != nil {
		return &pkg.GenericError{Msg: fmt.Sprintf("Error while reading response body: %v", readErr)}
	}
	if len(tmp) > 0 {
		unmarshalErr := json.Unmarshal(tmp, response)
		if unmarshalErr != nil {
			return &pkg.GenericError{Msg: fmt.Sprintf("Error while unmarshaling response: %v", unmarshalErr)}
		}
	}
	return nil
}

func (c *CspmClient) buildBaseUrl() string {
	return fmt.Sprintf("%v://%v", c.baseClient.Schema, c.baseClient.BaseUrl)
}
