package cwpp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
)

const apiVersion = "22.01"

func Test_BuildBaseUrlSuccessful(t *testing.T) {
	client, err := NewCwppClient(&ClientOptions{
		ConsoleUrl: "foo",
		ApiVersion: apiVersion,
		MaxRetries: 0,
		Schema:     "https",
	})
	assert.Nil(t, err)

	expectedBaseUrl := "foo/api/v22.01"
	assert.Equal(t, expectedBaseUrl, client.consoleUrl)
}

func Test_BuildBaseUrlInvalidAPIVersion(t *testing.T) {
	client, err := NewCwppClient(&ClientOptions{
		ConsoleUrl: "foo",
		ApiVersion: "boo",
		MaxRetries: 0,
		Schema:     "http",
	})
	assert.Nil(t, client)
	assert.Error(t, &internal.GenericError{}, err)
}
