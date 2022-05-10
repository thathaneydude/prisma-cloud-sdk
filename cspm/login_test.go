package cspm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"prisma-cloud-sdk/client"
	"testing"
)

func TestCspmClient_LoginFullRequest(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(loginEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(cspmClient.ContentTypeHeader, client.ApplicationJSON)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"customerNames":[{"customerName":"PANW","prismaId":"4321","tosAccepted":true}],"message":"foo","roles":["admin"],"token":"12345"}`))
	})

	loginResponse, err := client.Login("foo", "bar")
	assert.Nil(t, err)
	assert.Equal(t, loginResponse.Token, "12345")
}

func TestCspmClient_LoginInvalidCredentials(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%v", loginEndpoint), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client2.ContentTypeHeader, client2.ApplicationJSON)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Invalid Credentials`))
	})

	loginResponse, err := client.Login("foo", "bar")
	assert.Nil(t, loginResponse)
	assert.Error(t, &client2.UnauthorizedError{}, err)
}

func TestCspmClient_LoginInternalServerError(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%v", loginEndpoint), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client2.ContentTypeHeader, client2.ApplicationJSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Login Failed Unknown Error`))
	})

	loginResponse, err := client.Login("foo", "bar")
	assert.Nil(t, loginResponse)
	assert.Error(t, &client2.InternalServerError{}, err)
}
