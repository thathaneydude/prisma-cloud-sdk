package cspm

import "github.com/thathaneydude/prisma-cloud-sdk/internal"

const resourceEndpoint = "/resource"

func (c *CspmClient) GetResource(rrn string) (*Resource, error) {
	var resource Resource
	err := c.postWithResponseInterface(resourceEndpoint, internal.ToBytes(&ResourceRequest{Rrn: rrn}), &resource)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type ResourceRequest struct {
	Rrn string `json:"rrn"`
}

type Resource struct {
	Rrn                      string                   `json:"rrn"`
	Id                       string                   `json:"id"`
	Name                     string                   `json:"name"`
	Url                      string                   `json:"url"`
	AccountId                string                   `json:"accountId"`
	AccountName              string                   `json:"accountName"`
	CloudType                string                   `json:"cloudType"`
	RegionId                 string                   `json:"regionId"`
	RegionName               string                   `json:"regionName"`
	Service                  string                   `json:"service"`
	ResourceType             string                   `json:"resourceType"`
	InsertTs                 int64                    `json:"insertTs"`
	Deleted                  bool                     `json:"deleted"`
	VpcId                    string                   `json:"vpcId"`
	VpcName                  string                   `json:"vpcName"`
	Tags                     map[string]interface{}   `json:"tags"`
	RiskGrade                string                   `json:"riskGrade"`
	Data                     []map[string]interface{} `json:"data"`
	HasNetwork               bool                     `json:"hasNetwork"`
	HasExternalFinding       bool                     `json:"hasExternalFinding"`
	HasExternalIntegration   bool                     `json:"hasExternalIntegration"`
	AllowDrillDown           bool                     `json:"allowDrillDown"`
	HasExtFindingRiskFactors bool                     `json:"hasExtFindingRiskFactors"`
}
