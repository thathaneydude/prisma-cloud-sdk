package cwpp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	bc "github.com/thathaneydude/prisma-cloud-sdk/internal/client"
	"net/http"
	"testing"
)

const apiVersion = "22.01"

func TestCwppClient_Authenticate(t *testing.T) {
	teardown := setup()
	defer teardown()

	client, err := NewCwppClient(&ClientOptions{
		ConsoleUrl: server.URL,
		ApiVersion: apiVersion,
		MaxRetries: 0,
		Schema:     "http",
	})
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
	assert.Equal(t, "Bearer 12345", client.baseClient.Headers.Get(authHeader))

}

func Test_BuildBaseUrlSuccessful(t *testing.T) {
	client, err := NewCwppClient(&ClientOptions{
		ConsoleUrl: "foo",
		ApiVersion: apiVersion,
		MaxRetries: 0,
		Schema:     "https",
	})
	assert.Nil(t, err)

	expectedBaseUrl := "foo/api/v22.01"
	assert.Equal(t, expectedBaseUrl, client.consoleUrl)
}

func Test_BuildBaseUrlInvalidAPIVersion(t *testing.T) {
	client, err := NewCwppClient(&ClientOptions{
		ConsoleUrl: "foo",
		ApiVersion: "boo",
		MaxRetries: 0,
		Schema:     "http",
	})
	assert.Nil(t, client)
	assert.Error(t, &internal.GenericError{}, err)
}
