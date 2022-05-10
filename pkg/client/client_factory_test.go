package client

import (
	"PrismaCloud/pkg/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBaseClient_HTTP(t *testing.T) {
	baseUrl := "foo"
	schema := "http"

	c := NewBaseClient(false, constants.DefaultMaxRetries, "http")
	assert.Equal(t, schema, c.schema)
	assert.Equal(t, constants.DefaultMaxRetries, c.maxRetries)
	assert.NotNil(t, c.httpClient)
}

func TestNewBaseClient_HTTPS(t *testing.T) {
	baseUrl := "foo"
	schema := "https"

	c := NewBaseClient(false, constants.DefaultMaxRetries, "https")
	assert.Equal(t, schema, c.schema)
	assert.Equal(t, constants.DefaultMaxRetries, c.maxRetries)
	assert.NotNil(t, c.httpClient)
}

//TODO: Test for TLS config
