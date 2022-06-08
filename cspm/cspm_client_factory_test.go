package cspm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCSPMClient_Success(t *testing.T) {
	c, err := NewCSPMClient(&ClientOptions{
		ApiUrl:     "api.prismacloud.io",
		SslVerify:  false,
		Schema:     "http",
		MaxRetries: 3,
	})
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "api.prismacloud.io", c.baseUrl)
}

func TestNewCSPMClient_InvalidAPI(t *testing.T) {
	c, err := NewCSPMClient(&ClientOptions{
		ApiUrl:     "foo",
		SslVerify:  false,
		Schema:     "https",
		MaxRetries: 3,
	})
	assert.Nil(t, c)
	assert.Error(t, err)
}

func TestNewCSPMClient_LocalAPIValidationSkip(t *testing.T) {
	c, err := NewCSPMClient(&ClientOptions{
		ApiUrl:     "foo",
		SslVerify:  false,
		Schema:     "http",
		MaxRetries: 3,
	})
	assert.Nil(t, err)
	assert.Equal(t, "foo", c.baseUrl)
}
