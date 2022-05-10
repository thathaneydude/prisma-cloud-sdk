package internal

import (
	"net/http"
	"net/url"
)

type BaseClient interface {
	Do(req http.Request) (*http.Response, error)
	DoWithRetry(req http.Request, currentAttempt int) (*http.Response, error)
	BuildRequest(baseUrl string, method string, endpoint string, params url.Values, data []byte) (*http.Request, error)
}

type Client interface {
	GetWithResponseInterface(endpoint string, params url.Values, response interface{}) error
	PostWithResponseInterface(endpoint string, body []byte, response interface{}) error
	PutWithResponseInterface(endpoint string, body []byte, response interface{}) error
	PatchWithResponseInterface(endpoint string, body []byte, response interface{}) error
	DeleteWithResponseInterface(endpoint string, params url.Values, response interface{}) error
	Get(endpoint string, params url.Values) (*http.Response, error)
	Post(endpoint string, body []byte) (*http.Response, error)
	Put(endpoint string, body []byte) (*http.Response, error)
	Patch(endpoint string, body []byte) (*http.Response, error)
	Delete(endpoint string, params url.Values) (*http.Response, error)
}
