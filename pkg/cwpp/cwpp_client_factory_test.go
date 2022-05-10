package cwpp

import (
	"PrismaCloud/pkg"
	bc "PrismaCloud/pkg/client"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCwppClient_Authenticate(t *testing.T) {
	teardown := setup()
	defer teardown()

	apiVersion := "22.01"

	client, err := NewCwppClient(server.URL, apiVersion, false, "http")
	assert.Nil(t, err)

	fullUri := fmt.Sprintf("/api/v%v%v", apiVersion, authenticateEndpoint)

	mux.HandleFunc(fullUri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(authHeader, "foo")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"token": "12345"}`))
	})

	resp, err := client.Authenticate("palo", "alto")
	assert.Nil(t, err)
	assert.Equal(t, "12345", resp.Token)
	assert.Equal(t, "Bearer 12345", client.BaseClient.Headers.Get(authHeader))

}

func Test_BuildBaseUrlSuccessful(t *testing.T) {
	client, err := NewCwppClient("foo", "22.01", false, "http")
	assert.Nil(t, err)

	expectedBaseUrl := "foo/api/v22.01"
	assert.Equal(t, expectedBaseUrl, client.BaseClient.BaseUrl)
}

func Test_BuildBaseUrlInvalidAPIVersion(t *testing.T) {
	client, err := NewCwppClient("foo", "bar", false, "http")
	assert.Nil(t, client)
	assert.Error(t, &pkg.GenericError{}, err)
}
