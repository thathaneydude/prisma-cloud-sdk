package cspm

import (
	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	bc "github.com/thathaneydude/prisma-cloud-sdk/internal/client"
	"net/http"
	"testing"
)

func TestCspmClient_ExtendAuthTokenSuccessful(t *testing.T) {
	teardown := setup()
	defer teardown()
	cspmClient, err := NewCSPMClient(&ClientOptions{
		ApiUrl:     server.URL,
		SslVerify:  false,
		Schema:     "http",
		MaxRetries: 3,
	})
	assert.Nil(t, err)
	mux.HandleFunc(authExtendEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(internal.AuthHeader, "foo")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"customerNames":[{"customerName":"PANW","prismaId":"4321","tosAccepted":true}],"message":"foo","roles":["admin"],"token":"12345"}`))
	})

	loginResponse, err := cspmClient.ExtendAuthToken()
	assert.Nil(t, err)
	assert.Equal(t, loginResponse.Token, "12345")
}

func TestCspmClient_ExtendAuthTokenInternalServerError(t *testing.T) {
	teardown := setup()
	defer teardown()
	cspmClient, err := NewCSPMClient(&ClientOptions{
		ApiUrl:     server.URL,
		SslVerify:  false,
		Schema:     "http",
		MaxRetries: 3,
	})
	assert.Nil(t, err)
	mux.HandleFunc(authExtendEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(internal.AuthHeader, "foo")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`login_failed_unknown_error`))
	})

	loginResponse, err := cspmClient.ExtendAuthToken()
	assert.Nil(t, loginResponse)
	assert.Error(t, &internal.GenericError{}, err)
}
