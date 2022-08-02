package cwpp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/client"
)

func TestCwppClient_ListImages(t *testing.T) {

	fileContent, _ := ioutil.ReadFile("./test_files/list_images_response.json")

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
	fullUri := fmt.Sprintf("/api/v%v%v", apiVersion, imagesEndpoint)
	mux.HandleFunc(fullUri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client.ContentTypeHeader, client.ApplicationJSON)
		w.WriteHeader(http.StatusOK)
		w.Write(fileContent)
	})

	listResponse, err := c.ListImages(ImageQuery{
		Offset: "",
		Limit:  "",
	})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(listResponse))
	assert.Equal(t, "testname1", listResponse[0].Id)
	assert.Equal(t, "bash", listResponse[0].Binaries[0].Name)
	assert.Equal(t, "alpine", listResponse[1].OsDistro)

}
