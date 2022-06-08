package cspm

import (
	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	bc "github.com/thathaneydude/prisma-cloud-sdk/internal/client"
	"net/http"
	"testing"
)

func TestCspmClient_GetMetaInfo(t *testing.T) {
	teardown := setup()
	defer teardown()
	cspmClient, err := NewCSPMClient(&ClientOptions{
		ApiUrl:     server.URL,
		SslVerify:  false,
		Schema:     "http",
		MaxRetries: 3,
	})
	assert.Nil(t, err)
	mux.HandleFunc(metaInfoEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(internal.AuthHeader, "foo")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"licenseType":"enterprise","marketplace":"Palo Alto Networks Marketplace","startTs":1539027762524,"endTs":1667174400000,"twistlockUrl":"https://us-east1.cloud.twistlock.com/12345"}`))
	})

	metaInfo, err := cspmClient.GetMetaInfo()
	assert.Nil(t, err)
	assert.Equal(t, "enterprise", metaInfo.LicenseType)
	assert.Equal(t, "Palo Alto Networks Marketplace", metaInfo.Marketplace)
	assert.Equal(t, "https://us-east1.cloud.twistlock.com/12345", metaInfo.TwistlockUrl)
}

func TestCspmClient_GetMetaInfoFailsWithError(t *testing.T) {
	teardown := setup()
	defer teardown()
	cspmClient, err := NewCSPMClient(&ClientOptions{
		ApiUrl:     server.URL,
		SslVerify:  false,
		Schema:     "http",
		MaxRetries: 3,
	})
	assert.Nil(t, err)
	mux.HandleFunc(metaInfoEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(bc.ContentTypeHeader, bc.ApplicationJSON)
		w.Header().Set(internal.AuthHeader, "foo")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`internal server error`))
	})

	metaInfo, err := cspmClient.GetMetaInfo()
	assert.Nil(t, metaInfo)
	assert.Error(t, err)
}
