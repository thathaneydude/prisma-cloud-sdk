package cs

import "github.com/thathaneydude/prisma-cloud-sdk/cspm"

// NewDefaultCSClient is a pass-through to the CSPM client. Once created, client.Login(...) will need to be
// done to get the JWT and load it into the base client headers.
func NewDefaultCSClient(apiUrl string, sslVerify bool, schema string, maxRetries int) (*CsClient, error) {
	// Just a pass-through to CSPM
	CspmClient, err := cspm.NewCSPMClient(apiUrl, sslVerify, schema, maxRetries)
	if err != nil {
		return nil, err
	}
	return &CsClient{cspmClient: CspmClient}, nil
}

func CsClientWithCspmInjected(cspmClient *cspm.CspmClient) *CsClient {
	return &CsClient{cspmClient: cspmClient}
}
