package cwpp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/client"
)

func TestCwppClient_ListDefenders(t *testing.T) {

	fileContent, _ := ioutil.ReadFile("./test_files/list_defenders_response.json")

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
	fullUri := fmt.Sprintf("/api/v%v%v", apiVersion, defendedEndpoint)
	mux.HandleFunc(fullUri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client.ContentTypeHeader, client.ApplicationJSON)
		w.WriteHeader(http.StatusOK)
		w.Write(fileContent)
	})

	listResponse, err := c.ListDefenders(DefenderQuery{
		Offset: "",
		Limit:  "",
	})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(listResponse))
	assert.Equal(t, "testname1", listResponse[0].Hostname)
	assert.Equal(t, "aws", listResponse[1].CloudMetadata.Provider)
	assert.Equal(t, "test2", listResponse[1].Fqdn)

}
