package cwpp

import (
	"PrismaCloud/pkg/utils"
	"fmt"
)

const (
	authHeader           = "Authorization"
	authenticateEndpoint = "/authenticate"
)

func (c *CwppClient) Authenticate(username string, password string) (*AuthenticateResponse, error) {
	c.username = username
	c.password = password
	authRequest := AuthenticateRequest{
		Username: username,
		Password: password,
	}

	var authResponse AuthenticateResponse
	err := c.PostWithResponseInterface(authenticateEndpoint, utils.ToBytes(authRequest), &authResponse)
	if err != nil {
		return nil, err
	}
	c.baseClient.Headers.Set(authHeader, fmt.Sprintf("Bearer %v", authResponse.Token))

	return &authResponse, nil
}

type AuthenticateRequest struct {
	Password string `json:"password"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}
