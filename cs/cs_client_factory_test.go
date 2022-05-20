package cs

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	logrus.SetOutput(ioutil.Discard)

	return func() {
		server.Close()
	}
}

// Putting the mock server setup here since all the tests are done in the CSPM package. No real factory tests needed.
