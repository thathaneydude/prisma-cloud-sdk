package cwpp

import (
	"fmt"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	bc "github.com/thathaneydude/prisma-cloud-sdk/internal/client"
)

type ClientOptions struct {
	ConsoleUrl string
	ApiVersion string
	SslVerify  bool
	MaxRetries int
	Schema     string
}

// NewCwppClient creates a CwppClient from the ClientOptions provided. The client will need to Authenticate first
// before making any requests
func NewCwppClient(o *ClientOptions) (*CwppClient, error) {
	cwppBaseUrl, err := buildBaseUrl(o.ConsoleUrl, o.ApiVersion)
	if err != nil {
		return nil, err
	}

	baseClient := bc.NewBaseClient(o.SslVerify, o.MaxRetries, o.Schema)
	return &CwppClient{
		baseClient: *baseClient,
		consoleUrl: cwppBaseUrl,
		apiVersion: o.ApiVersion,
	}, nil
}

func buildBaseUrl(baseUrl string, apiVersion string) (string, error) {
	validatedVersion, err := validateApiVersion(apiVersion)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v/api/v%v", baseUrl, validatedVersion), nil
}

func validateApiVersion(apiVersion string) (string, error) {
	found := false
	for _, allowedApiVersion := range internal.APIVersions {
		if allowedApiVersion == apiVersion {
			found = true
			break
		}
	}
	if found == false {
		return "", &internal.GenericError{Msg: fmt.Sprintf("API version \"%v\" provided is not a valid option: %v", apiVersion, internal.APIVersions)}
	}
	return apiVersion, nil
}
