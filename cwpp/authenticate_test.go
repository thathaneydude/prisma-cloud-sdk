package cwpp

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	bc "github.com/thathaneydude/prisma-cloud-sdk/internal/client"
)

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

	resp, err := client.Authenticate("", "")

	assert.Nil(t, err)
	assert.NotEqual(t, "23456", resp.Token)
	assert.Equal(t, "Bearer 12345", client.baseClient.Headers.Get(authHeader))

	//TODO: Fix authentication failure error handling
}
