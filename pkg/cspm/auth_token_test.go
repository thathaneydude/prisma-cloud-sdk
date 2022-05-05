package cspm

import (
	"PrismaCloud/pkg"
	bc "PrismaCloud/pkg/client"
	"github.com/stretchr/testify/assert"
	"net/http"
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

	loginResponse, err := client.ExtendAuthToken()
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

	loginResponse, err := client.ExtendAuthToken()
	assert.Nil(t, loginResponse)
	assert.Error(t, &pkg.GenericError{}, err)
}
