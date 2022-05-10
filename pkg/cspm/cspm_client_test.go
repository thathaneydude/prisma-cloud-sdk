package cspm

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *CspmClient
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client, _ = NewCSPMClient(server.URL, false, "http")
	logrus.SetOutput(ioutil.Discard)

	return func() {
		server.Close()
	}
}
