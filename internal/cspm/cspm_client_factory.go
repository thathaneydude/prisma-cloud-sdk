package cspm

import (
	"fmt"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	bc "github.com/thathaneydude/prisma-cloud-sdk/internal/client"
	"golang.org/x/exp/slices"
)

type ClientOptions struct {
	ApiUrl     string
	SslVerify  bool
	Schema     string
	MaxRetries int
}

// NewCSPMClient should really only be used by the
func NewCSPMClient(o *ClientOptions) (*CspmClient, error) {
	if o.Schema == "https" && !slices.Contains(internal.SupportedAPIURLs, o.ApiUrl) {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("API url provided \"%v\" is not supported. Please reference %v for more information", o.ApiUrl, internal.SupportedAPIURLLink)}
	}

	baseClient := bc.NewBaseClient(o.SslVerify, o.MaxRetries, o.Schema)
	return &CspmClient{
		BaseClient: *baseClient,
		baseUrl:    o.ApiUrl,
	}, nil
}
