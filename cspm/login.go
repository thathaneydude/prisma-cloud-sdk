package cspm

import (
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
)

const loginEndpoint = "/login"

// Login Returns a JWT auth token for accessing the Prisma Cloud APIs
//
// https://prisma.pan.dev/api/cloud/cspm/login#operation/app-login
func (c *CspmClient) Login(loginReq *LoginRequest) (*LoginResponse, error) {
	var loginResponse LoginResponse
	err := c.postWithResponseInterface(loginEndpoint, internal.ToBytes(loginReq), &loginResponse)
	if err != nil {
		return nil, &internal.GenericError{Msg: err.Error()}
	}

	c.baseClient.Headers.Set(internal.AuthHeader, loginResponse.Token)
	return &loginResponse, nil
}

type LoginRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CustomerName string `json:"customerName,omitempty"`
	PrismaId     string `json:"prismaId,omitempty"`
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
