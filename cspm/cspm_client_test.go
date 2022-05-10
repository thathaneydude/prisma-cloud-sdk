package cspm

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var (
	mux        *http.ServeMux
	server     *httptest.Server
	cspmClient *CspmClient
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	cspmClient, _ = NewCSPMClient(server.URL, false, "http", 3)
	logrus.SetOutput(ioutil.Discard)

	return func() {
		server.Close()
	}
}
