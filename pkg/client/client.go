package client

import (
	"PrismaCloud/pkg"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const sleepDuration = 5

type BaseClientImpl struct {
	httpClient *http.Client
	Headers    *http.Header
	BaseUrl    string
	Schema     string
	maxRetries int
}

func (c BaseClientImpl) Do(req http.Request) (*http.Response, error) {
	req.Header = *c.Headers
	resp, err := c.httpClient.Do(&req)
	if err != nil {
		logrus.Errorf("Failed to make request: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c BaseClientImpl) DoWithRetry(req http.Request, currentAttempt int) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

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
		respBody, _ := ioutil.ReadAll(resp.Body)
		return nil, &UnauthorizedError{msg: string(respBody)}
	case http.StatusInternalServerError:
		respBody, _ := ioutil.ReadAll(resp.Body)
		return nil, &InternalServerError{msg: string(respBody)}
	case http.StatusNotFound:
		respBody, _ := ioutil.ReadAll(resp.Body)
		return nil, &NotFoundError{msg: string(respBody)}
	}
	return resp, nil
}

func (c BaseClientImpl) BuildRequest(method string, endpoint string, params url.Values, data []byte) (*http.Request, error) {
	if !slices.Contains([]string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete}, method) {
		return nil, &pkg.GenericError{Msg: fmt.Sprintf("Improper HTTP method provided: %v", method)}
	}

	base, err := url.Parse(c.BaseUrl)
	if err != nil {
		return nil, &pkg.GenericError{Msg: fmt.Sprintf("Error while parsing base url provided \"%v\": %v", c.BaseUrl, err)}
	}

	if base.Scheme == "" {
		base.Scheme = c.Schema
	}
	// Require client func to provide separating "/"
	base.Path += endpoint

	var payload bytes.Buffer
	if data != nil {
		jsonData, marshalErr := json.Marshal(data)
		if marshalErr != nil {
			return nil, &pkg.GenericError{Msg: fmt.Sprintf("Error while converting provided body to JSON: %v", marshalErr)}
		}
		payload = *bytes.NewBuffer(jsonData)
	}

	if params != nil {
		base.RawQuery = params.Encode()
	}

	req, reqErr := http.NewRequest(method, base.String(), &payload)

	if reqErr != nil {
		return nil, &pkg.GenericError{Msg: fmt.Sprintf("Error while building request: %v", reqErr)}
	}
	logrus.Debugf("Request: %v", req)
	return req, nil
}
