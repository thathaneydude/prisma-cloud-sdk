package cspm

import (
	"fmt"
	bc "github.com/thathaneydude/prisma-cloud-sdk/client"
	"github.com/thathaneydude/prisma-cloud-sdk/constants"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"golang.org/x/exp/slices"
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
