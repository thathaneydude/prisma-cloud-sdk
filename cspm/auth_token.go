package cspm

import (
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"github.com/thathaneydude/prisma-cloud-sdk/utils"
	"net/http"
)

const authExtendEndpoint = "/auth_token/extend"

func (c *CspmClient) ExtendAuthToken() (*LoginResponse, error) {
	// ExtendAuthToken LoginResponse is the same response for extending the token
	var authTokenExtendResponse LoginResponse
	req, err := c.BaseClient.BuildRequest(c.baseUrl, http.MethodGet, authExtendEndpoint, nil, nil)
	resp, err := c.BaseClient.Do(*req)
	if err != nil {
		return nil, &internal.GenericError{Msg: err.Error()}
	}
	err = utils.UnmarshalResponse(resp, &authTokenExtendResponse)
	if err != nil {
		return nil, err
	}
	return &authTokenExtendResponse, nil
}
