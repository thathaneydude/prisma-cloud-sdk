package cwpp

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"net/http"
)

const (
	authHeader           = "Authorization"
	authenticateEndpoint = "/authenticate"
)

// Authenticate Requests a JWT from Twistlock with the username and password provided. API Keys should be used if possible.
// If the request is successful, the Authorization Bearer token will be updated with the response
//
// https://prisma.pan.dev/api/cloud/cwpp/authenticate#operation/post-authenticate
func (c *CwppClient) Authenticate(username string, password string) (*AuthenticateResponse, error) {
	c.username = username
	c.password = password
	authRequest := AuthenticateRequest{
		Username: username,
		Password: password,
	}

	var authResponse AuthenticateResponse
	marshalledRequest, err := json.Marshal(authRequest)
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to marshal login request body: %v", err)}
	}
	req, err := c.baseClient.BuildRequest(c.consoleUrl, http.MethodPost, authenticateEndpoint, nil, marshalledRequest)
	resp, err := c.baseClient.Do(*req)
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to authenticate: %v", err.Error())}
	}
	err = internal.UnmarshalResponse(resp, &authResponse)
	if err != nil {
		return nil, err
	}
	c.baseClient.Headers.Set(authHeader, fmt.Sprintf("Bearer %v", authResponse.Token))
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
