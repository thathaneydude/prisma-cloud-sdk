package client

import (
	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"testing"
)

func TestNewBaseClient_HTTP(t *testing.T) {
	schema := "http"

	c := NewBaseClient(false, internal.DefaultMaxRetries, "http")
	assert.Equal(t, schema, c.schema)
	assert.Equal(t, internal.DefaultMaxRetries, c.maxRetries)
	assert.NotNil(t, c.httpClient)
}

func TestNewBaseClient_HTTPS(t *testing.T) {
	schema := "https"

	c := NewBaseClient(false, internal.DefaultMaxRetries, "https")
	assert.Equal(t, schema, c.schema)
	assert.Equal(t, internal.DefaultMaxRetries, c.maxRetries)
	assert.NotNil(t, c.httpClient)
}

//TODO: Test for TLS config
