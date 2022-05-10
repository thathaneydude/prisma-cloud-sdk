package cspm

import (
	"fmt"
	"golang.org/x/exp/slices"
	"prisma-cloud-sdk/pkg"
	bc "prisma-cloud-sdk/pkg/client"
	"prisma-cloud-sdk/pkg/constants"
)

func NewCSPMClient(apiUrl string, sslVerify bool, schema string, maxRetries int) (*CspmClient, error) {
	if !slices.Contains(constants.SupportedAPIURLs, apiUrl) {
		return nil, &pkg.GenericError{Msg: fmt.Sprintf("API url provided \"%v\" is not supported. Please reference %v for more information", apiUrl, constants.SupportedAPIURLLink)}
	}

	baseClient := bc.NewBaseClient(sslVerify, maxRetries, schema)
	return &CspmClient{
		BaseClient: *baseClient,
		baseUrl:    apiUrl,
	}, nil
}
