package cspm

import (
	"encoding/json"
	"fmt"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"net/url"
)

const accountGroupsEndpoint = "/cloud/group"

// ListAccountGroups Returns an array of accessible account groups
//
// https://prisma.pan.dev/api/cloud/cspm/account-groups#operation/get-account-groups
func (c *CspmClient) ListAccountGroups(excludeCloudAccountDetails bool) ([]AccountGroupResponse, error) {
	var accountGroups []AccountGroupResponse
	params := url.Values{}
	params.Set("excludeCloudAccountDetails", fmt.Sprintf("%v", excludeCloudAccountDetails))
	err := c.getWithResponseInterface(accountGroupsEndpoint, params, &accountGroups)
	if err != nil {
		return nil, err
	}
	return accountGroups, nil
}

// AddAccountGroup Create a new account group on the Prisma Cloud platform specifying the attributes in
// an AccountGroup
//
// https://prisma.pan.dev/api/cloud/cspm/account-groups#operation/add-account-group
func (c *CspmClient) AddAccountGroup(accountGroup AccountGroup) (*AccountGroupResponse, error) {
	var accountGroupResp AccountGroupResponse
	marshalledRequest, err := json.Marshal(accountGroup)
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to marshal request body: %v", err)}
	}
	err = c.postWithResponseInterface(accountGroupsEndpoint, marshalledRequest, &accountGroupResp)
	if err != nil {
		return nil, err
	}
	return &accountGroupResp, nil
}

// UpdateAccountGroup Update information related to an existing account group with the attributes of the provided
// AccountGroup
//
// https://prisma.pan.dev/api/cloud/cspm/account-groups#operation/update-account-group
func (c *CspmClient) UpdateAccountGroup(accountGroupId string, newAccountGroup AccountGroup) error {
	marshalledRequest, err := json.Marshal(newAccountGroup)
	if err != nil {
		return &internal.GenericError{Msg: fmt.Sprintf("Failed to marshal request body: %v", err)}
	}
	_, err = c.Put(fmt.Sprintf("%v/%v", accountGroupsEndpoint, accountGroupId), marshalledRequest)
	if err != nil {
		return err
	}
	return nil
}

type AccountGroup struct {
	Name        string   `json:"name"`
	AccountIds  []string `json:"accountIds"`
	Description string   `json:"description,omitempty"`
}

type AccountGroupResponse struct {
	AccountIds []string `json:"accountIds,omitempty"`
	Accounts   []struct {
		Id   string `json:"id,omitempty"`
		Name string `json:"name"`
		Type string `json:"type,omitempty"`
	} `json:"accounts,omitempty"`
	AlertRules []struct {
		AlertId   string `json:"alertId,omitempty"`
		AlertName string `json:"alertName,omitempty"`
	} `json:"alertRules,omitempty"`
	AutoCreated       bool `json:"autoCreated,omitempty"`
	CloudAccountCount int  `json:"cloudAccountCount,omitempty"`
	CloudAccountInfos []struct {
		AccountId      string `json:"accountId,omitempty"`
		CloudType      string `json:"cloudType,omitempty"`
		LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	} `json:"cloudAccountInfos,omitempty"`
	Description    string `json:"description,omitempty"`
	Id             string `json:"id,omitempty"`
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	LastModifiedTs int    `json:"lastModifiedTs,omitempty"`
	Name           string `json:"name,omitempty"`
	ParentInfo     struct {
		AutoCreated bool   `json:"autoCreated,omitempty"`
		Id          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
	} `json:"parentInfo,omitempty"`
}
