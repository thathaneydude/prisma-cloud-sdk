package cwpp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/client"
)

func TestCwppClient_ListContainers(t *testing.T) {

	fileContent, _ := ioutil.ReadFile("./test_files/list_containers_response.json")

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
	fullUri := fmt.Sprintf("/api/v%v%v", apiVersion, containersEndpoint)
	mux.HandleFunc(fullUri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client.ContentTypeHeader, client.ApplicationJSON)
		w.WriteHeader(http.StatusOK)
		w.Write(fileContent)
	})

	listResponse, err := c.ListContainers(ContainerQuery{
		Offset: "0",
		Limit:  "50",
	})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(listResponse))
	assert.Equal(t, "testid1", listResponse[0].ID)
	assert.Equal(t, "testname1", listResponse[0].Info.Name)
	assert.Equal(t, 11, listResponse[1].Info.ComplianceIssuesCount)
	assert.Equal(t, "annotation.io.kubernetes.pod.terminationGracePeriod:30", listResponse[1].Info.Labels[0])
}
