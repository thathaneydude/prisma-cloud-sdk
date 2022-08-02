package cwpp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/sirupsen/logrus"
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
