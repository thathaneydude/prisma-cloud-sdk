package cspm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	bc "prisma-cloud-sdk/pkg/client"
	"testing"
)

func TestCspmClient_LoginFullRequest(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(loginEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"customerNames":[{"customerName":"PANW","prismaId":"4321","tosAccepted":true}],"message":"foo","roles":["admin"],"token":"12345"}`))
	})

	loginResponse, err := client.Login("foo", "bar", "PANW", "4321")
	assert.Nil(t, err)
	assert.Equal(t, loginResponse.Token, "12345")
}

func TestCspmClient_LoginInvalidCredentials(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%v", loginEndpoint), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Invalid Credentials`))
	})

	loginResponse, err := client.Login("foo", "bar", "PANW", "4321")
	assert.Nil(t, loginResponse)
	assert.Error(t, &bc.UnauthorizedError{}, err)
}

func TestCspmClient_LoginInternalServerError(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/%v", loginEndpoint), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Login Failed Unknown Error`))
	})

	loginResponse, err := client.Login("foo", "bar", "PANW", "4321")
	assert.Nil(t, loginResponse)
	assert.Error(t, &bc.InternalServerError{}, err)
}
