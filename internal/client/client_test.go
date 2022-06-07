package client

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *BaseClient
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewBaseClient(false, 3, "http")
	logrus.SetOutput(ioutil.Discard)

	return func() {
		server.Close()
	}
}

func TestBaseClientImpl_DoWithRetry200(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(ContentTypeHeader, ApplicationJSON)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`blah`))
	})

	req, err := client.BuildRequest(server.URL, "GET", "foo", nil, nil)
	assert.Nil(t, err)
	resp, err := client.DoWithRetry(*req, 1)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestBaseClientImpl_DoWithRetry404(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(ContentTypeHeader, ApplicationJSON)
		w.WriteHeader(http.StatusNotFound)
	})

	req, err := client.BuildRequest(server.URL, "GET", "foo", nil, nil)
	assert.Nil(t, err)
	resp, err := client.DoWithRetry(*req, 1)
	assert.Equal(t, &NotFoundError{}, err)
	assert.Nil(t, resp)
}

func TestBaseClientImpl_DoWithRetry500(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(ContentTypeHeader, ApplicationJSON)
		w.WriteHeader(http.StatusInternalServerError)
	})

	req, err := client.BuildRequest(server.URL, "GET", "foo", nil, nil)
	assert.Nil(t, err)
	resp, err := client.DoWithRetry(*req, 1)
	assert.Equal(t, &InternalServerError{}, err)
	assert.Nil(t, resp)
}

func TestBaseClientImpl_DoWithRetry401(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(ContentTypeHeader, ApplicationJSON)
		w.WriteHeader(http.StatusUnauthorized)
	})

	req, err := client.BuildRequest(server.URL, "GET", "foo", nil, nil)
	assert.Nil(t, err)
	resp, err := client.DoWithRetry(*req, 1)
	assert.Equal(t, &UnauthorizedError{}, err)
	assert.Nil(t, resp)
}
