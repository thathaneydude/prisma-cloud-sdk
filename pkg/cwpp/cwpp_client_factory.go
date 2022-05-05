package cwpp

import (
	"PrismaCloud/pkg"
	bc "PrismaCloud/pkg/client"
	"PrismaCloud/pkg/constants"
	"fmt"
	"golang.org/x/exp/slices"
)

func NewCwppClient(baseUrl string, apiVersion string, sslVerify bool, schema string) (*CwppClient, error) {
	cwppBaseUrl, err := buildBaseUrl(baseUrl, apiVersion)
	if err != nil {
		return nil, err
	}
	baseClient := bc.NewBaseClient(cwppBaseUrl, sslVerify, constants.DefaultMaxRetries, schema)
	return &CwppClient{baseClient: *baseClient, apiVersion: apiVersion}, nil
}

func NewCwppClientWithCustomMaxRetries(baseUrl string, apiVersion string, sslVerify bool, maxRetries int, schema string) (*CwppClient, error) {
	cwppBaseUrl, err := buildBaseUrl(baseUrl, apiVersion)
	if err != nil {
		return nil, err
	}
	baseClient := bc.NewBaseClient(cwppBaseUrl, sslVerify, maxRetries, schema)
	return &CwppClient{baseClient: *baseClient}, nil
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
