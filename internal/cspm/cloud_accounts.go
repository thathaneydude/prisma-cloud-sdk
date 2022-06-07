package cspm

import (
	"fmt"
	"net/url"
)

const cloudAccountEndpoint = "/cloud"

func (c *CspmClient) ListCloudAccounts(excludeAccountGroupDetails bool) ([]CloudAccountResponse, error) {
	var cloudAccounts []CloudAccountResponse
	params := url.Values{}
	params.Set("excludeAccountGroupDetails", fmt.Sprintf("%v", excludeAccountGroupDetails))
	err := c.GetWithResponseInterface(cloudAccountEndpoint, params, &cloudAccounts)
	if err != nil {
		return nil, err
	}
	return cloudAccounts, nil
}

type CloudAccountResponse struct {
	Name               string   `json:"name,omitempty"`
	CloudType          string   `json:"cloudType,omitempty"`
	AccountType        string   `json:"accountType,omitempty"`
	Enabled            bool     `json:"enabled,omitempty"`
	LastModifiedTs     int64    `json:"lastModifiedTs,omitempty"`
	LastModifiedBy     string   `json:"lastModifiedBy,omitempty"`
	StorageScanEnabled bool     `json:"storageScanEnabled,omitempty"`
	ProtectionMode     string   `json:"protectionMode,omitempty"`
	IngestionMode      int      `json:"ingestionMode,omitempty"`
	DeploymentType     string   `json:"deploymentType,omitempty"`
	GroupIds           []string `json:"groupIds,omitempty"`
	Groups             []struct {
		Id   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"groups,omitempty"`
	Status                string `json:"status,omitempty"`
	NumberOfChildAccounts int    `json:"numberOfChildAccounts,omitempty"`
	AccountId             string `json:"accountId,omitempty"`
	AddedOn               int64  `json:"addedOn,omitempty"`
}
