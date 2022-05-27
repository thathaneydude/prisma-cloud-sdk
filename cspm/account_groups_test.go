package cspm

import (
	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/client"
	"net/http"
	"testing"
)

func TestCspmClient_ListAccountGroups(t *testing.T) {
	teardown := setup()
	defer teardown()
	c, err := NewCSPMClient(server.URL, false, "http", 3)
	assert.Nil(t, err)
	mux.HandleFunc(accountGroupsEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(client.ContentTypeHeader, client.ApplicationJSON)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":"db99166d-e096-4b03-819e-423b7dfe7fcc","name":"foo","description":"test group","lastModifiedBy":"foo@paloaltonetworks.com","lastModifiedTs":1653501619490,"accountIds":[],"nonOnboardedCloudAccountIds":[],"autoCreated":false,"accounts":[],"alertRules":[],"parentInfo":{}},{"id":"d4d00dd1-297d-4a03-8783-0648816b8a71","name":"bar","description":"","lastModifiedBy":"bar@paloaltonetworks.com","lastModifiedTs":1653499831455,"accountIds":[],"nonOnboardedCloudAccountIds":[],"autoCreated":false,"accounts":[],"alertRules":[{"alertId":"9d0a5306-3029-4291-8825-ac60bb1cd91d","alertName":"bar Test"}],"parentInfo":{}},{"id":"659d0468-cc1d-4f9c-9bb2-5b85f3135a9e","name":"baz","description":"baz test","lastModifiedBy":"baz@paloaltonetworks.com","lastModifiedTs":1653499775180,"accountIds":["12345678"],"nonOnboardedCloudAccountIds":[],"autoCreated":false,"accounts":[{"id":"12345678","name":"baz-green","type":"aws"},{"id":"87654321","name":"baz-blue","type":"gcp"}],"alertRules":[],"parentInfo":{}}]`))
	})

	listResponse, err := c.ListAccountGroups(false)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(listResponse))
	assert.Equal(t, "foo", listResponse[0].Name)
	assert.Equal(t, "test group", listResponse[0].Description)
	assert.Equal(t, "bar", listResponse[1].Name)
	assert.Equal(t, "", listResponse[1].Description)
	assert.Equal(t, 1, len(listResponse[1].AlertRules))
	assert.Equal(t, "9d0a5306-3029-4291-8825-ac60bb1cd91d", listResponse[1].AlertRules[0].AlertId)
	assert.Equal(t, "bar Test", listResponse[1].AlertRules[0].AlertName)
	assert.Equal(t, "baz", listResponse[2].Name)
	assert.Equal(t, "baz test", listResponse[2].Description)
	assert.Equal(t, 1, len(listResponse[2].AccountIds))
	assert.Equal(t, "12345678", listResponse[2].AccountIds[0])
	assert.Equal(t, 2, len(listResponse[2].Accounts))
	assert.Equal(t, "12345678", listResponse[2].Accounts[0].Id)
	assert.Equal(t, "baz-green", listResponse[2].Accounts[0].Name)
	assert.Equal(t, "aws", listResponse[2].Accounts[0].Type)
	assert.Equal(t, "87654321", listResponse[2].Accounts[1].Id)
	assert.Equal(t, "baz-blue", listResponse[2].Accounts[1].Name)
	assert.Equal(t, "gcp", listResponse[2].Accounts[1].Type)

}
