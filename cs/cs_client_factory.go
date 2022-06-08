package cs

import (
	cspm2 "github.com/thathaneydude/prisma-cloud-sdk/cspm"
)

// NewDefaultCSClient is a pass-through to the CSPM client. Once created, client.Login(...) will need to be
// done to get the JWT and load it into the base client headers.
func NewDefaultCSClient(o *cspm2.ClientOptions) (*CsClient, error) {
	// Just a pass-through to CSPM
	CspmClient, err := cspm2.NewCSPMClient(o)
	if err != nil {
		return nil, err
	}
	return &CsClient{cspmClient: CspmClient}, nil
}

func CsClientWithCspmInjected(cspmClient *cspm2.CspmClient) *CsClient {
	return &CsClient{cspmClient: cspmClient}
}
