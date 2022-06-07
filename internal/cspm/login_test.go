package cspm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/client"
	"net/http"
	"testing"
)

func TestCspmClient_LoginFullRequest(t *testing.T) {
	teardown := setup()
	defer teardown()
	cspmClient, err := NewCSPMClient(&ClientOptions{
		ApiUrl:     server.URL,
		SslVerify:  false,
		Schema:     "http",
		MaxRetries: 3,
	})
	mux.HandleFunc(loginEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client.ContentTypeHeader, client.ApplicationJSON)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"customerNames":[{"customerName":"PANW","prismaId":"4321","tosAccepted":true}],"message":"foo","roles":["admin"],"token":"12345"}`))
	})

	loginResponse, err := cspmClient.Login(nil)
	assert.Nil(t, err)
	assert.Equal(t, loginResponse.Token, "12345")
}

func TestCspmClient_LoginInvalidCredentials(t *testing.T) {
	teardown := setup()
	defer teardown()
	cspmClient, err := NewCSPMClient(&ClientOptions{
		ApiUrl:     server.URL,
		SslVerify:  false,
		Schema:     "http",
		MaxRetries: 3,
	})
	mux.HandleFunc(fmt.Sprintf("/%v", loginEndpoint), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client.ContentTypeHeader, client.ApplicationJSON)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Invalid Credentials`))
	})

	loginResponse, err := cspmClient.Login(nil)
	assert.Nil(t, loginResponse)
	assert.Error(t, &client.UnauthorizedError{}, err)
}

func TestCspmClient_LoginInternalServerError(t *testing.T) {
	teardown := setup()
	defer teardown()
	cspmClient, err := NewCSPMClient(&ClientOptions{
		ApiUrl:     server.URL,
		SslVerify:  false,
		Schema:     "http",
		MaxRetries: 3,
	})
	mux.HandleFunc(fmt.Sprintf("/%v", loginEndpoint), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client.ContentTypeHeader, client.ApplicationJSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Login Failed Unknown Error`))
	})

	loginResponse, err := cspmClient.Login(nil)
	assert.Nil(t, loginResponse)
	assert.Error(t, &client.InternalServerError{}, err)
}
