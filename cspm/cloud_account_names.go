package cspm

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"net/url"
)

const cloudAccountNamesEndpoint = "/cloud/name"

func (c *CspmClient) ListCloudAccountNames(query ListCloudAccountNamesQuery) ([]CloudAccountResponse, error) {
	var encoder = schema.NewEncoder()
	params := url.Values{}
	err := encoder.Encode(query, params)
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to decode query provided: %v", err)}
	}
	var cloudAccountNames []CloudAccountResponse
	err = c.GetWithResponseInterface(cloudAccountNamesEndpoint, params, &cloudAccountNames)
	if err != nil {
		return nil, err
	}
	return cloudAccountNames, nil
}

type ListCloudAccountNamesQuery struct {
	OnlyActive     bool     `schema:"onlyActive,omitempty"`
	AmountGroupIds []string `schema:"accountGroupIds,omitempty"`
	CloudType      string   `schema:"cloudType,omitempty"`
}

type CloudAccountNamesResponse struct {
	CloudType         string `json:"cloudType,omitempty"`
	Id                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	ParentAccountName string `json:"parentAccountName,omitempty"`
}
