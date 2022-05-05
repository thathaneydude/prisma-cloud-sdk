package cspm

import (
	bc "PrismaCloud/pkg/client"
)

const MaxRetries = 3

func NewCSPMClient(apiUrl string, sslVerify bool) CspmClient {
	baseClient := bc.NewBaseClient(apiUrl, sslVerify, MaxRetries)
	return CspmClient{baseClient: *baseClient}
}

func NewCSPMClientWithCustomMaxRetries(apiUrl string, sslVerify bool, maxRetries int) CspmClient {
	baseClient := bc.NewBaseClient(apiUrl, sslVerify, maxRetries)
	return CspmClient{baseClient: *baseClient}
}
