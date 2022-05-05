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
	maxRetries int
}

func (c BaseClientImpl) Do(req http.Request) (*http.Response, error) {
	req.Header = *c.Headers
	resp, err := c.httpClient.Do(&req)
	if err != nil {
		logrus.Errorf("Failed to make request: %v", err)
		return nil, err
	}
	respText, err := ioutil.ReadAll(resp.Body)
	logrus.Debugf("Response: %v", string(respText))
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
			resp, err = c.DoWithRetry(req, currentAttempt+1)
		} else {
			logrus.Errorf("Maximum number of retry attempts (%v) exceeded", c.maxRetries)
		}
	case http.StatusUnauthorized:
		respBody, _ := ioutil.ReadAll(resp.Body)
		return nil, &UnauthorizedError{msg: string(respBody)}
	case http.StatusInternalServerError:
		respBody, _ := ioutil.ReadAll(resp.Body)
		return nil, &InternalServerError{msg: string(respBody)}
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

	req, reqErr := http.NewRequest(method, base.RequestURI(), &payload)

	if reqErr != nil {
		return nil, &pkg.GenericError{Msg: fmt.Sprintf("Error while building request: %v", reqErr)}
	}
	logrus.Debugf("Request: %v", req)
	return req, nil
}

//
//func (c BaseClientImpl) Get(endpoint string, params string) (*http.Response, error) {
//	var fullUrl string
//	if params != "" {
//		fullUrl = fmt.Sprintf("%v/%v?%v", c.BaseUrl, endpoint, params)
//	} else {
//		fullUrl = fmt.Sprintf("%v/%v", c.BaseUrl, endpoint)
//	}
//	logrus.Debugf("GET --> %v", fullUrl)
//
//	req, err := http.NewRequest("GET", fullUrl, nil)
//	if err != nil {
//		logrus.Errorf("Unable to build GET request \"%v\": %v", fullUrl, err)
//		return nil, err
//	}
//	resp, err := c.DoWithRetry(*req, 1)
//	if err != nil {
//		logrus.Errorf("Error while running GET request \"%v\": %v", fullUrl, err)
//		return nil, err
//	}
//	return resp, nil
//}
//
//func (c BaseClientImpl) Post(endpoint string, body []byte) (*http.Response, error) {
//	fullUrl := fmt.Sprintf("%v/%v", c.BaseUrl, endpoint)
//	logrus.Debugf("POST -> %v : %v", fullUrl, string(body))
//
//	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
//	if err != nil {
//		logrus.Errorf("Unable to build POST request \"%v\": %v", fullUrl, err)
//		return nil, err
//	}
//
//	resp, err := c.DoWithRetry(*req, 1)
//	if err != nil {
//		logrus.Errorf("Error while running POST request %v: %v", fullUrl, err)
//		return nil, err
//	}
//	return resp, err
//}
//
//func (c BaseClientImpl) Put(endpoint string, body []byte) (*http.Response, error) {
//	fullUrl := fmt.Sprintf("%v/%v", c.BaseUrl, endpoint)
//	log.Printf("PUT -> %v : %v", fullUrl, string(body))
//	req, err := http.NewRequest("PUT", fullUrl, bytes.NewBuffer(body))
//	if err != nil {
//		logrus.Errorf("Unable to build PUT request \"%v\": %v", fullUrl, err)
//		return nil, err
//	}
//
//	resp, err := c.DoWithRetry(*req, 1)
//	if err != nil {
//		logrus.Errorf("Error while running PUT request \"%v\": %v", fullUrl, err)
//		return nil, err
//	}
//	return resp, err
//}
//
//func (c BaseClientImpl) Patch(endpoint string, body []byte) (*http.Response, error) {
//	fullUrl := fmt.Sprintf("%v/%v", c.BaseUrl, endpoint)
//	log.Printf("PATCH -> %v : %v", fullUrl, string(body))
//	req, err := http.NewRequest("PATCH", fullUrl, bytes.NewBuffer(body))
//	if err != nil {
//		logrus.Errorf("Unable to build PATCH request \"%v\": %v", fullUrl, err)
//		return nil, err
//	}
//
//	resp, err := c.DoWithRetry(*req, 1)
//	if err != nil {
//		logrus.Errorf("Error while running PATCH request \"%v\": %v", fullUrl, err)
//		return nil, err
//	}
//	return resp, err
//}
//
//func (c BaseClientImpl) Delete(endpoint string, params string) (*http.Response, error) {
//	var fullUrl string
//	if params != "" {
//		fullUrl = fmt.Sprintf("%v/%v?%v", c.BaseUrl, endpoint, params)
//	} else {
//		fullUrl = fmt.Sprintf("%v/%v", c.BaseUrl, endpoint)
//	}
//	logrus.Debugf("DELETE --> %v", fullUrl)
//
//	req, err := http.NewRequest("DELETE", fullUrl, nil)
//	if err != nil {
//		logrus.Errorf("Unable to build DELETE request \"%v\": %v", fullUrl, err)
//		return nil, err
//	}
//	resp, err := c.DoWithRetry(*req, 1)
//	if err != nil {
//		logrus.Errorf("Error while running DELETE request \"%v\": %v", fullUrl, err)
//		return nil, err
//	}
//	return resp, nil
//}
