package cspm

import (
	"fmt"
	"github.com/thathaneydude/prisma-cloud-sdk/utils"
	"net/url"
)

const accountGroupsEndpoint = "/cloud/group"

func (c *CspmClient) ListAccountGroups(excludeCloudAccountDetails bool) ([]AccountGroupResponse, error) {
	var accountGroups []AccountGroupResponse
	params := url.Values{}
	params.Set("excludeCloudAccountDetails", fmt.Sprintf("%v", excludeCloudAccountDetails))
	err := c.GetWithResponseInterface(accountGroupsEndpoint, params, &accountGroups)
	if err != nil {
		return nil, err
	}
	return accountGroups, nil
}

func (c *CspmClient) AddAccountGroup(accountGroup AccountGroup) (*AccountGroupResponse, error) {
	var accountGroupResp AccountGroupResponse
	err := c.PostWithResponseInterface(accountGroupsEndpoint, utils.ToBytes(accountGroup), &accountGroupResp)
	if err != nil {
		return nil, err
	}
	return &accountGroupResp, nil
}

func NewAccountGroup(name string, accountIds []string) AccountGroup {
	return newAccountGroup(name, accountIds, "")
}

func NewAccountGroupWithDescription(name string, accountIds []string, description string) AccountGroup {
	return newAccountGroup(name, accountIds, description)
}

func newAccountGroup(name string, accountIds []string, description string) AccountGroup {
	ag := &AccountGroup{
		Name:       name,
		AccountIds: accountIds,
	}

	if description != "" {
		ag.Description = description
	}

	return *ag
}

type AccountGroup struct {
	Name        string   `schema:"name"`
	AccountIds  []string `schema:"accountIds"`
	Description string   `schema:"description,omitempty"`
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
