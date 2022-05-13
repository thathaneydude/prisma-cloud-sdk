package cspm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCSPMClient_Success(t *testing.T) {
	c, err := NewCSPMClient("api.prismacloud.io", false, "https", 3)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "api.prismacloud.io", c.baseUrl)
}

func TestNewCSPMClient_InvalidAPI(t *testing.T) {
	c, err := NewCSPMClient("foo", false, "https", 3)
	assert.Nil(t, c)
	assert.Error(t, err)
}

func TestNewCSPMClient_LocalAPIValidationSkip(t *testing.T) {
	c, err := NewCSPMClient("foo", false, "http", 3)
	assert.Nil(t, err)
	assert.Equal(t, "foo", c.baseUrl)
}
