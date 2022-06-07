package cs

import (
	"github.com/thathaneydude/prisma-cloud-sdk/internal/cspm"
)

func (c *CsClient) Login(username string, password string) (*cspm.LoginResponse, error) {
	return c.cspmClient.Login(username, password)
}
