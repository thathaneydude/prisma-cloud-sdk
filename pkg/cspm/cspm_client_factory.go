package cspm

import (
	bc "PrismaCloud/pkg/client"
	"PrismaCloud/pkg/constants"
)

func NewCSPMClient(apiUrl string, sslVerify bool, schema string) CspmClient {
	baseClient := bc.NewBaseClient(apiUrl, sslVerify, constants.DefaultMaxRetries, schema)
	return CspmClient{baseClient: *baseClient}
}

func NewCSPMClientWithCustomMaxRetries(apiUrl string, sslVerify bool, maxRetries int, schema string) CspmClient {
	baseClient := bc.NewBaseClient(apiUrl, sslVerify, maxRetries, schema)
	return CspmClient{baseClient: *baseClient}
}
