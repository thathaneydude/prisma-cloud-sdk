package client

import (
	"PrismaCloud/pkg/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBaseClient_HTTP(t *testing.T) {
	baseUrl := "foo"
	schema := "http"

	c := NewBaseClient(baseUrl, false, constants.DefaultMaxRetries, "http")
	assert.Equal(t, baseUrl, c.BaseUrl)
	assert.Equal(t, schema, c.Schema)
	assert.Equal(t, constants.DefaultMaxRetries, c.maxRetries)
	assert.NotNil(t, c.httpClient)
}

func TestNewBaseClient_HTTPS(t *testing.T) {
	baseUrl := "foo"
	schema := "https"

	c := NewBaseClient(baseUrl, false, constants.DefaultMaxRetries, "https")
	assert.Equal(t, baseUrl, c.BaseUrl)
	assert.Equal(t, schema, c.Schema)
	assert.Equal(t, constants.DefaultMaxRetries, c.maxRetries)
	assert.NotNil(t, c.httpClient)
}

//TODO: Test for TLS config
