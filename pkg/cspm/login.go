package cspm

import (
	"PrismaCloud/pkg"
	"PrismaCloud/pkg/utils"
)

const (
	loginEndpoint = "/login"
	authHeader    = "x-redlock-auth"
)

func (c *CspmClient) Login(username string, password string, customerName string, prismaId string) (*LoginResponse, error) {
	loginRequest := LoginRequest{
		CustomerName: customerName,
		Password:     password,
		PrismaID:     prismaId,
		Username:     username,
	}
	var loginResponse LoginResponse
	err := c.PostWithResponseInterface(loginEndpoint, utils.ToBytes(loginRequest), &loginResponse)
	if err != nil {
		return nil, &pkg.GenericError{Msg: err.Error()}
	}

	c.baseClient.Headers.Set(authHeader, loginResponse.Token)
	return &loginResponse, nil
}

type LoginRequest struct {
	CustomerName string `json:"customerName,omitempty"`
	Password     string `json:"password"`
	PrismaID     string `json:"prismaId,omitempty"`
	Username     string `json:"username"`
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
