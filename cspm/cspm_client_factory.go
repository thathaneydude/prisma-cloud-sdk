package cspm

import (
	"fmt"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	bc "github.com/thathaneydude/prisma-cloud-sdk/internal/client"
)

type ClientOptions struct {
	ApiUrl     string
	SslVerify  bool
	Schema     string
	MaxRetries int
}

func NewCSPMClient(o *ClientOptions) (*CspmClient, error) {
	if o.Schema == "https" {
		found := false
		for _, supportedApiUrl := range internal.SupportedAPIURLs {
			if o.ApiUrl == supportedApiUrl {
				found = true
				break
			}
		}
		if found == false {
			return nil, &internal.GenericError{Msg: fmt.Sprintf("API url provided \"%v\" is not supported. Please reference %v for more information", o.ApiUrl, internal.SupportedAPIURLLink)}
		}
	}

	baseClient := bc.NewBaseClient(o.SslVerify, o.MaxRetries, o.Schema)
	return &CspmClient{
		baseClient: *baseClient,
		baseUrl:    o.ApiUrl,
	}, nil
}
