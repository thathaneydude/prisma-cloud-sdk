package cspm

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"golang.org/x/exp/slices"
	"net/url"
)

const cloudAccountNamesEndpoint = "/cloud/name"

// ListCloudAccountNames Returns a list of cloud account IDs and names based on the provided
// ListCloudAccountNamesQuery
//
// https://prisma.pan.dev/api/cloud/cspm/cloud-accounts#operation/get-cloud-account-names
func (c *CspmClient) ListCloudAccountNames(query ListCloudAccountNamesQuery) ([]CloudAccountResponse, error) {
	var encoder = schema.NewEncoder()
	params := url.Values{}
	err := encoder.Encode(query, params)
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to decode query provided: %v", err)}
	}
	var cloudAccountNames []CloudAccountResponse
	err = c.getWithResponseInterface(cloudAccountNamesEndpoint, params, &cloudAccountNames)
	if err != nil {
		return nil, err
	}
	return cloudAccountNames, nil
}

// NewListCloudAccountNamesQuery creates a query for use with ListCloudAccountNames
func NewListCloudAccountNamesQuery(onlyActive bool, amountGroupIds []string, cloudType string) (*ListCloudAccountNamesQuery, error) {
	if !slices.Contains(internal.CloudTypes, cloudType) {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Cloud type %v provided is not supported. Must be one of the following: %v", cloudType, internal.CloudTypes)}
	}
	return &ListCloudAccountNamesQuery{
		OnlyActive:     onlyActive,
		AmountGroupIds: amountGroupIds,
		CloudType:      cloudType,
	}, nil
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
