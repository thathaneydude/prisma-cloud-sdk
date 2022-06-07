package cspm

import (
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
)

const listUserV3Endpoint = "/v3/user"

func (c *CspmClient) ListUsersV3() ([]UserV3, error) {
	var userListV3 []UserV3
	err := c.GetWithResponseInterface(listUserV3Endpoint, nil, &userListV3)
	if err != nil {
		return nil, err
	}
	return userListV3, nil
}

func (c *CspmClient) AddUserV3(req AddUserV3Request) (*AddUserV3Response, error) {
	var addUserV3 AddUserV3Response
	err := c.PostWithResponseInterface(listUserV3Endpoint, internal.ToBytes(req), &addUserV3)
	if err != nil {
		return nil, err
	}
	return &addUserV3, nil
}

type AddUserV3Request struct {
	AccessKeyExpiration int      `json:"accessKeyExpiration,omitempty"`
	AccessKeyName       string   `json:"accessKeyName,omitempty"`
	AccessKeysAllowed   bool     `json:"accessKeysAllowed,omitempty"`
	DefaultRoleId       string   `json:"defaultRoleId,omitempty"`
	Email               string   `json:"email,omitempty"`
	EnableKeyExpiration bool     `json:"enableKeyExpiration,omitempty"`
	FirstName           string   `json:"firstName,omitempty"`
	LastName            string   `json:"lastName,omitempty"`
	RoleIds             []string `json:"roleIds,omitempty"`
	TimeZone            string   `json:"timeZone,omitempty"`
	Type                string   `json:"type,omitempty"`
	Username            string   `json:"username,omitempty"`
}

type AddUserV3Response struct {
	Id        string `json:"id,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`
}

type UserV3 struct {
	AccessKeyExpiration int    `json:"accessKeyExpiration"`
	AccessKeyName       string `json:"accessKeyName"`
	AccessKeysAllowed   bool   `json:"accessKeysAllowed"`
	AccessKeysCount     int    `json:"accessKeysCount"`
	ActiveRole          struct {
		Id                     string `json:"id"`
		Name                   string `json:"name"`
		OnlyAllowCIAccess      bool   `json:"onlyAllowCIAccess"`
		OnlyAllowComputeAccess bool   `json:"onlyAllowComputeAccess"`
		OnlyAllowReadAccess    bool   `json:"onlyAllowReadAccess"`
		Type                   string `json:"type"`
	} `json:"activeRole"`
	DefaultRoleId       string   `json:"defaultRoleId"`
	DisplayName         string   `json:"displayName"`
	Email               string   `json:"email"`
	EnableKeyExpiration bool     `json:"enableKeyExpiration"`
	Enabled             bool     `json:"enabled"`
	FirstName           string   `json:"firstName"`
	LastLoginTs         int      `json:"lastLoginTs"`
	LastModifiedBy      string   `json:"lastModifiedBy"`
	LastModifiedTs      int      `json:"lastModifiedTs"`
	LastName            string   `json:"lastName"`
	RoleIds             []string `json:"roleIds"`
	Roles               []struct {
		Id                     string `json:"id"`
		Name                   string `json:"name"`
		OnlyAllowCIAccess      bool   `json:"onlyAllowCIAccess"`
		OnlyAllowComputeAccess bool   `json:"onlyAllowComputeAccess"`
		OnlyAllowReadAccess    bool   `json:"onlyAllowReadAccess"`
		Type                   string `json:"type"`
	} `json:"roles"`
	TimeZone string `json:"timeZone"`
	Type     string `json:"type"`
	Username string `json:"username"`
}
