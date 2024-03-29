package cwpp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/client"
)

func TestCwppClient_ListCollections(t *testing.T) {

	fileContent, _ := ioutil.ReadFile("./test_files/list_collections_response.json")

	teardown := setup()
	defer teardown()
	c, err := NewCwppClient(&ClientOptions{
		ConsoleUrl: server.URL,
		SslVerify:  false,
		ApiVersion: apiVersion,
		Schema:     "http",
		MaxRetries: 3,
	})

	assert.Nil(t, err)
	fullUri := fmt.Sprintf("/api/v%v%v", apiVersion, collectionEndpoint)
	mux.HandleFunc(fullUri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client.ContentTypeHeader, client.ApplicationJSON)
		w.WriteHeader(http.StatusOK)
		w.Write(fileContent)
	})

	listResponse, err := c.ListCollections()

	assert.Nil(t, err)
	assert.Equal(t, 3, len(listResponse))
	assert.Equal(t, "test1", listResponse[0].Name)
	assert.Equal(t, "System - cloud account test3 collection", listResponse[2].Description)
	assert.Equal(t, "testid2", listResponse[1].AccountIds[0])

}
