package cspm

import (
	"fmt"
	"golang.org/x/exp/slices"
	bc "prisma-cloud-sdk/client"
	"prisma-cloud-sdk/constants"
	"prisma-cloud-sdk/internal"
)

func NewCSPMClient(apiUrl string, sslVerify bool, schema string, maxRetries int) (*CspmClient, error) {
	if !slices.Contains(constants.SupportedAPIURLs, apiUrl) {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("API url provided \"%v\" is not supported. Please reference %v for more information", apiUrl, constants.SupportedAPIURLLink)}
	}

	baseClient := bc.NewBaseClient(sslVerify, maxRetries, schema)
	return &CspmClient{
		BaseClient: *baseClient,
		baseUrl:    apiUrl,
	}, nil
}
