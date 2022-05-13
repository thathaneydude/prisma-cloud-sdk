package cwpp

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"github.com/thathaneydude/prisma-cloud-sdk/utils"
	"net/http"
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
	req, err := c.BaseClient.BuildRequest(c.consoleUrl, http.MethodPost, authenticateEndpoint, nil, utils.ToBytes(authRequest))
	resp, err := c.BaseClient.Do(*req)
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to authenticate: %v", err.Error())}
	}
	err = utils.UnmarshalResponse(resp, &authResponse)
	if err != nil {
		return nil, err
	}
	c.BaseClient.Headers.Set(authHeader, fmt.Sprintf("Bearer %v", authResponse.Token))
	logrus.Debugf("Setting %v header to %v", authHeader, authResponse.Token)
	return &authResponse, nil
}

type AuthenticateRequest struct {
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
	Username string `json:"username"`
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}
