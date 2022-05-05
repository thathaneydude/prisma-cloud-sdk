package cspm

import (
	"PrismaCloud/pkg"
)

const authExtendEndpoint = "/auth_token/extend"

func (c *CspmClient) ExtendAuthToken() (*LoginResponse, error) {
	// ExtendAuthToken LoginResponse is the same response for extending the token
	var authTokenExtendResponse LoginResponse
	err := c.GetWithResponseInterface(authExtendEndpoint, nil, authTokenExtendResponse)
	if err != nil {
		return nil, &pkg.GenericError{Msg: err.Error()}
	}
	return &authTokenExtendResponse, nil
}
