package cspm

import (
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
)

const loginEndpoint = "/login"

func (c *CspmClient) Login(username string, password string) (*LoginResponse, error) {
	loginRequest := LoginRequest{
		Username: username,
		Password: password,
	}
	var loginResponse LoginResponse
	err := c.PostWithResponseInterface(loginEndpoint, internal.ToBytes(loginRequest), &loginResponse)
	if err != nil {
		return nil, &internal.GenericError{Msg: err.Error()}
	}

	c.BaseClient.Headers.Set(internal.AuthHeader, loginResponse.Token)
	return &loginResponse, nil
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	CustomerNames []struct {
		CustomerName string `json:"customerName,omitempty"`
		PrismaID     string `json:"prismaId,omitempty"`
		TosAccepted  bool   `json:"tosAccepted,omitempty"`
	} `json:"customerNames,omitempty"`
	Message string   `json:"message,omitempty"`
	Roles   []string `json:"roles,omitempty"`
	Token   string   `json:"token"`
}
