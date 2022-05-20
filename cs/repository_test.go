package cs

import (
	"github.com/stretchr/testify/assert"
	bc "github.com/thathaneydude/prisma-cloud-sdk/client"
	"github.com/thathaneydude/prisma-cloud-sdk/constants"
	"net/http"
	"testing"
)

func TestCsClient_ListRepositories(t *testing.T) {
	teardown := setup()
	defer teardown()
	csClient, err := NewDefaultCSClient(server.URL, false, "http", 3)
	assert.Nil(t, err)

	mux.HandleFunc(listRepositoriesEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(constants.AuthHeader, "foo")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":"f87d47af-fe09-4a86-a0e5-05e0d11ff9b8","repository":"aws","source":"cli","owner":"foo","defaultBranch":null,"isPublic":false,"runs":1,"creationDate":"2022-03-14T09:47:14.301Z","lastScanDate":null},{"id":"8d2307ac-f5af-4cc8-9b86-34a57581174c","repository":"scripts","source":"Github","owner":"bar","defaultBranch":"main","isPublic":false,"runs":0,"creationDate":"2022-04-12T06:59:32.694Z","lastScanDate":"2022-05-18T13:25:46.016Z"},{"id":"4621352c-d732-41b6-988c-b491cf581440","repository":"1","source":"cli","owner":"baz","defaultBranch":null,"isPublic":false,"runs":0,"creationDate":"2022-02-19T11:53:08.351Z","lastScanDate":null}]`))
	})

	repos, err := csClient.ListRepositories(false)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(repos))
	assert.Equal(t, "f87d47af-fe09-4a86-a0e5-05e0d11ff9b8", repos[0].Id)
	assert.Equal(t, "aws", repos[0].Repository)
	assert.Equal(t, "cli", repos[0].Source)
	assert.Equal(t, "foo", repos[0].Owner)
	assert.Equal(t, "", repos[0].DefaultBranch)
	assert.Equal(t, false, repos[0].IsPublic)
	assert.Equal(t, 1, repos[0].Runs)
	assert.Equal(t, "", repos[0].LastScanDate)
}
