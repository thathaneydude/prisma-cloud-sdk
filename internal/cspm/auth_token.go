package cspm

import (
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"net/http"
)

const authExtendEndpoint = "/auth_token/extend"

// ExtendAuthToken attempts to extend the JWT from a previous Login (Prisma Cloud)
//
// https://prisma.pan.dev/api/cloud/cspm/login#operation/extend-session
func (c *CspmClient) ExtendAuthToken() (*LoginResponse, error) {
	// ExtendAuthToken LoginResponse is the same response for extending the token
	var authTokenExtendResponse LoginResponse
	req, err := c.baseClient.BuildRequest(c.baseUrl, http.MethodGet, authExtendEndpoint, nil, nil)
	resp, err := c.baseClient.Do(*req)
	if err != nil {
		return nil, &internal.GenericError{Msg: err.Error()}
	}
	err = internal.UnmarshalResponse(resp, &authTokenExtendResponse)
	if err != nil {
		return nil, err
	}
	return &authTokenExtendResponse, nil
}
