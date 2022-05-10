package cwpp

import (
	"PrismaCloud/pkg"
	bc "PrismaCloud/pkg/client"
	"PrismaCloud/pkg/constants"
	"fmt"
	"golang.org/x/exp/slices"
)

func NewCwppClient(consoleUrl string, apiVersion string, sslVerify bool, schema string) (*CwppClient, error) {
	cwppBaseUrl, err := buildBaseUrl(consoleUrl, apiVersion)
	if err != nil {
		return nil, err
	}

	baseClient := bc.NewBaseClient(sslVerify, constants.DefaultMaxRetries, schema)
	return &CwppClient{
		BaseClient: *baseClient,
		consoleUrl: cwppBaseUrl,
		apiVersion: apiVersion,
	}, nil
}

func NewCwppClientWithCustomMaxRetries(consoleUrl string, apiVersion string, sslVerify bool, maxRetries int, schema string) (*CwppClient, error) {
	cwppBaseUrl, err := buildBaseUrl(consoleUrl, apiVersion)
	if err != nil {
		return nil, err
	}
	baseClient := bc.NewBaseClient(sslVerify, maxRetries, schema)
	return &CwppClient{
		BaseClient: *baseClient,
		consoleUrl: cwppBaseUrl,
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
	if !slices.Contains(constants.APIVersions, apiVersion) {
		return "", &pkg.GenericError{Msg: fmt.Sprintf("API version \"%v\" provided is not a valid option: %v", apiVersion, constants.APIVersions)}
	}
	return apiVersion, nil
}
