package prisma

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	bc "github.com/thathaneydude/prisma-cloud-sdk/client"
	"github.com/thathaneydude/prisma-cloud-sdk/constants"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestNewPrismaCloudClient(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(constants.AuthHeader, "foo")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"customerNames":[{"customerName":"PANW","prismaId":"4321","tosAccepted":true}],"message":"foo","roles":["admin"],"token":"12345"}`))
	})

	mux.HandleFunc("/meta_info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(constants.AuthHeader, "foo")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"licenseType":"enterprise","marketplace":"Palo Alto Networks Marketplace","startTs":1539027762524,"endTs":1667174400000,"twistlockUrl":"https://us-east1.cloud.twistlock.com/12345"}`))
	})

	c, err := NewPrismaCloudClient(server.URL, "http", "foo", "bar", "22.01", 3, false)
	assert.Nil(t, err)
	assert.NotNil(t, c)
}

func TestNewPrismaCloudClient_AuthFail(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(constants.AuthHeader, "foo")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`blah`))
	})

	mux.HandleFunc("/auth_token/extend", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(constants.AuthHeader, "foo")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`blah`))
	})

	mux.HandleFunc("/meta_info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(constants.AuthHeader, "foo")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"licenseType":"enterprise","marketplace":"Palo Alto Networks Marketplace","startTs":1539027762524,"endTs":1667174400000,"twistlockUrl":"https://us-east1.cloud.twistlock.com/12345"}`))
	})

	c, err := NewPrismaCloudClient(server.URL, "http", "foo", "bar", "22.01", 3, false)
	assert.Nil(t, c)
	assert.Error(t, &internal.GenericError{}, err)
}
