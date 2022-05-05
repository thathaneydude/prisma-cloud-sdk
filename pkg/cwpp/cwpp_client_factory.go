package cwpp

import (
	"PrismaCloud/pkg"
	bc "PrismaCloud/pkg/client"
	"fmt"
	"golang.org/x/exp/slices"
)

const MaxRetries = 3

func NewCwppClient(baseUrl string, apiVersion string, sslVerify bool) (*CwppClient, error) {
	fullBaseUrl, err := buildBaseUrl(baseUrl, apiVersion)
	if err != nil {
		return nil, err
	}
	baseClient := bc.NewBaseClient(fullBaseUrl, sslVerify, MaxRetries)
	return &CwppClient{baseClient: *baseClient, apiVersion: apiVersion}, nil
}

func NewCwppClientWithCustomMaxRetries(baseUrl string, apiVersion string, sslVerify bool, maxRetries int) (*CwppClient, error) {
	fullBaseUrl, err := buildBaseUrl(baseUrl, apiVersion)
	if err != nil {
		return nil, err
	}
	baseClient := bc.NewBaseClient(fullBaseUrl, sslVerify, maxRetries)
	return &CwppClient{baseClient: *baseClient}, nil
}

func buildBaseUrl(baseUrl string, apiVersion string) (string, error) {
	validatedVersion, err := validateApiVersion(apiVersion)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%v/api/v%v", baseUrl, validatedVersion), nil
}

func validateApiVersion(apiVersion string) (string, error) {
	APIVersions := []string{"22.01", "21.08", "21.04"}
	if !slices.Contains(APIVersions, apiVersion) {
		return "", &pkg.GenericError{Msg: fmt.Sprintf("API version \"%v\" provided is not a valid option: %v", apiVersion, APIVersions)}
	}
	return apiVersion, nil
}
