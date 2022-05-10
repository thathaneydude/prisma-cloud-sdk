package cspm

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	bc "prisma-cloud-sdk/client"
	"prisma-cloud-sdk/internal"
	"testing"
)

func TestCspmClient_ExtendAuthTokenSuccessful(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(authExtendEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(authHeader, "foo")
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

	mux.HandleFunc(authExtendEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(authHeader, "foo")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`login_failed_unknown_error`))
	})

	loginResponse, err := cspmClient.ExtendAuthToken()
	assert.Nil(t, loginResponse)
	assert.Error(t, &internal.GenericError{}, err)
}