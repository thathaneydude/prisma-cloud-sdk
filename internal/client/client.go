package client

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"golang.org/x/exp/slices"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	sleepDuration          = 5
	redLockRequestIdHeader = "x-redlock-request-id"
)

type BaseClient struct {
	httpClient *http.Client
	Headers    *http.Header
	schema     string
	maxRetries int
}

func (c *BaseClient) Do(req http.Request) (*http.Response, error) {
	req.Header = *c.Headers
	resp, err := c.httpClient.Do(&req)
	if err != nil {
		logrus.Errorf("Failed to make request: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *BaseClient) DoWithRetry(req http.Request, currentAttempt int) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	var respMsg string
	if resp.Header.Get(redLockRequestIdHeader) != "" {
		respMsg = fmt.Sprintf("Response: Status [%v]", resp.Status)
	} else {
		respMsg = fmt.Sprintf("Response: Status [%v] Request-ID [%v]", resp.Status, resp.Header.Get(redLockRequestIdHeader))
	}

	logrus.Debugf(respMsg)
	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		if currentAttempt < c.maxRetries {
			sleep := sleepDuration * time.Second
			logrus.Debugf("Waiting %v before retrying", sleep)
			return c.DoWithRetry(req, currentAttempt+1)
		} else {
			logrus.Errorf("Maximum number of retry attempts (%v) exceeded", c.maxRetries)
		}
	case http.StatusUnauthorized:
		// TODO: close response
		respBody, _ := ioutil.ReadAll(resp.Body)
		err = &UnauthorizedError{msg: string(respBody)}
		logrus.Errorf(err.Error())
		return nil, err
	case http.StatusInternalServerError:
		respBody, _ := ioutil.ReadAll(resp.Body)
		err = &InternalServerError{msg: string(respBody)}
		logrus.Errorf(err.Error())
		return nil, err
	case http.StatusNotFound:
		respBody, _ := ioutil.ReadAll(resp.Body)
		err = &NotFoundError{msg: string(respBody)}
		logrus.Errorf(err.Error())
		return nil, err
	case http.StatusMethodNotAllowed:
		respBody, _ := ioutil.ReadAll(resp.Body)
		err = &NotAllowedError{msg: string(respBody)}
		logrus.Errorf(err.Error())
		return nil, err
	}
	return resp, nil
}

func (c *BaseClient) BuildRequest(baseUrl string, method string, endpoint string, params url.Values, data []byte) (*http.Request, error) {
	if !slices.Contains(internal.SupportedHttpMethods, method) {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Improper HTTP method provided: %v", method)}
	}

	base, err := url.Parse(baseUrl)
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Error while parsing base url provided \"%v\": %v", baseUrl, err)}
	}

	if base.Scheme == "" {
		base.Scheme = c.schema
	}
	// Require client func to provide separating "/"
	base.Path += endpoint

	payload := bytes.NewBuffer(data)

	if params != nil {
		base.RawQuery = params.Encode()
	}

	req, reqErr := http.NewRequest(method, base.String(), payload)

	if reqErr != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Error while building request: %v", reqErr)}
	}
	logrus.Debugf("Request: %v", req)
	return req, nil
}
