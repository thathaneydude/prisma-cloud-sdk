package cwpp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/client"
)

func TestCwppClient_ListHosts(t *testing.T) {

	fileContent, _ := ioutil.ReadFile("./test_files/list_hosts_response.json")

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
	fullUri := fmt.Sprintf("/api/v%v%v", apiVersion, hostsEndpoint)
	mux.HandleFunc(fullUri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client.ContentTypeHeader, client.ApplicationJSON)
		w.WriteHeader(http.StatusOK)
		w.Write(fileContent)
	})

	listResponse, err := c.ListHosts(HostsQuery{
		Offset: "",
		Limit:  "",
	})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(listResponse))
	assert.Equal(t, "testname1", listResponse[0].Id)
	assert.Equal(t, "sshd", listResponse[1].Binaries[0].Name)
	assert.Equal(t, 17, listResponse[1].ComplianceDistribution.High)
	assert.Equal(t, false, listResponse[1].Agentless)

}
