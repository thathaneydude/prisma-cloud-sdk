package cspm

import (
	"fmt"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
)

const (
	listUserV3Endpoint = "/v3/user"
	userAccountType    = "USER_ACCOUNT"
	serviceAccountType = "SERVICE_ACCOUNT"
)

// ListUsersV3 Lists all users and service accounts for your tenant.
//
// https://prisma.pan.dev/api/cloud/cspm/user-profile#operation/get-user-profiles-v3
func (c *CspmClient) ListUsersV3() ([]UserV3, error) {
	var userListV3 []UserV3
	err := c.getWithResponseInterface(listUserV3Endpoint, nil, &userListV3)
	if err != nil {
		return nil, err
	}
	return userListV3, nil
}

// AddUserV3 Adds either a user profile or a service account profile
//
// https://prisma.pan.dev/api/cloud/cspm/user-profile#operation/add-user-v3
func (c *CspmClient) AddUserV3(req AddUserV3Request) (*AddUserV3Response, error) {
	var addUserV3 AddUserV3Response
	err := c.postWithResponseInterface(listUserV3Endpoint, internal.ToBytes(req), &addUserV3)
	if err != nil {
		return nil, err
	}
	return &addUserV3, nil
}

// NewUserAccountV3Request creates the AddUserV3Request needed when running AddUserV3 for a user account
func NewUserAccountV3Request(email string, firstName string, lastName string, roleIds []string, defaultRoleId string, timeZone string) (*AddUserV3Request, error) {
	if email == "" || firstName == "" || len(roleIds) == 0 || defaultRoleId == "" || timeZone == "" {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("All User Account parameters must be populated")}
	}
	return &AddUserV3Request{
		DefaultRoleId: defaultRoleId,
		Email:         email,
		FirstName:     firstName,
		LastName:      lastName,
		RoleIds:       roleIds,
		TimeZone:      timeZone,
		Type:          userAccountType,
	}, nil
}

// NewServiceAccountV3Request creates the AddUserV3Request needed when running AddUserV3 for a service account
func NewServiceAccountV3Request(username string, accessKeyName string, enableKeyExpiration bool, accessKeyExpiration int, defaultRoleId string, timeZone string) (*AddUserV3Request, error) {
	if username == "" || accessKeyName == "" || (enableKeyExpiration == true && accessKeyExpiration == 0) || defaultRoleId == "" || timeZone == "" {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("All Service Account parameters must be populated")}
	}

	req := &AddUserV3Request{
		AccessKeyName: accessKeyName,
		DefaultRoleId: defaultRoleId,
		TimeZone:      timeZone,
		Type:          serviceAccountType,
		Username:      username,
	}

	if enableKeyExpiration {
		req.EnableKeyExpiration = true
		req.AccessKeyExpiration = accessKeyExpiration
	}
	return req, nil
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
