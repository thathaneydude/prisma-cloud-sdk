package cwpp

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

type CwppClient struct {
	baseClient bc.BaseClientImpl
	apiVersion string
	username   string
	password   string
}

func (c *CwppClient) GetWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Get(endpoint, params)
	if err != nil {
		return err
	}
	return c.unmarshalResponse(resp, response)
}

func (c *CwppClient) PostWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Post(endpoint, body)
	if err != nil {
		return err
	}
	return c.unmarshalResponse(resp, response)
}

func (c *CwppClient) PutWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Put(endpoint, body)
	if err != nil {
		return err
	}
	return c.unmarshalResponse(resp, response)
}

func (c *CwppClient) PatchWithResponseInterface(endpoint string, body []byte, response interface{}) error {
	resp, err := c.Patch(endpoint, body)
	if err != nil {
		return err
	}
	return c.unmarshalResponse(resp, response)
}

func (c *CwppClient) DeleteWithResponseInterface(endpoint string, params url.Values, response interface{}) error {
	resp, err := c.Delete(endpoint, params)
	if err != nil {
		return err
	}
	return c.unmarshalResponse(resp, response)
}

func (c *CwppClient) Get(endpoint string, params url.Values) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest("GET", endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req, 1)
}

func (c *CwppClient) Post(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest("POST", endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req, 1)
}

func (c *CwppClient) Put(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest("PUT", endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req, 1)
}

func (c *CwppClient) Patch(endpoint string, body []byte) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest("PATCH", endpoint, nil, body)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req, 1)
}

func (c *CwppClient) Delete(endpoint string, params url.Values) (*http.Response, error) {
	req, err := c.baseClient.BuildRequest("DELETE", endpoint, params, nil)
	if err != nil {
		return nil, err
	}
	return c.cwppDoWithRetry(*req, 1)
}

func (c *CwppClient) cwppDoWithRetry(req http.Request, currentAttempt int) (*http.Response, error) {
	resp, err := c.baseClient.DoWithRetry(req, currentAttempt)
	sErr, _ := err.(*bc.UnauthorizedError)
	if sErr != nil {
		logrus.Debugf("Auth token may have expired. Attempting to refresh token")
		_, err = c.Authenticate(c.username, c.password)
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

func (c *CwppClient) unmarshalResponse(httpResponse *http.Response, response interface{}) error {
	defer httpResponse.Body.Close()
	tmp, readErr := io.ReadAll(httpResponse.Body)
	logrus.Debugf(string(tmp))
	if readErr != nil {
		return &pkg.GenericError{Msg: fmt.Sprintf("Error while reading response body: %v", readErr)}
	}
	unmarshalErr := json.Unmarshal(tmp, response)
	if unmarshalErr != nil {
		return &pkg.GenericError{Msg: fmt.Sprintf("Error while unmarshaling response: %v", unmarshalErr)}
	}
	return nil
}
