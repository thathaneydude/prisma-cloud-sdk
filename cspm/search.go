package cspm

import (
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
)

const searchConfigEndpoint = "/search/config"

func (c *CspmClient) SearchConfig(req SearchConfigRequest) (*SearchConfigResponse, error) {
	var searchResponse SearchConfigResponse
	err := c.postWithResponseInterface(searchConfigEndpoint, internal.ToBytes(req), &searchResponse)
	if err != nil {
		return nil, err
	}
	return &searchResponse, nil
}

type SearchConfigRequest struct {
	Id                string `json:"id,omitempty"`
	Limit             int    `json:"limit,omitempty"`
	Query             string `json:"query,omitempty"`
	SearchDescription string `json:"searchDescription,omitempty"`
	SearchName        string `json:"searchName,omitempty"`
	Sort              []struct {
		Direction string `json:"direction,omitempty"`
		Field     string `json:"field,omitempty"`
	} `json:"sort,omitempty"`
	TimeRange struct {
		RelativeTimeType string `json:"relativeTimeType,omitempty"`
		Type             string `json:"type,omitempty"`
		Value            struct {
			Amount int    `json:"amount,omitempty"`
			Unit   string `json:"unit,omitempty"`
		} `json:"value,omitempty"`
	} `json:"timeRange,omitempty"`
	WithResourceJson bool `json:"withResourceJson,omitempty"`
}

type SearchConfigResponse struct {
	CloudType   string `json:"cloudType"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SearchType  string `json:"searchType"`
	Saved       bool   `json:"saved"`
	TimeRange   struct {
		Type  string `json:"type"`
		Value struct {
			Unit   string `json:"unit"`
			Amount int    `json:"amount"`
		} `json:"value"`
		RelativeTimeType string `json:"relativeTimeType"`
	} `json:"timeRange"`
	Query string `json:"query"`
	Data  struct {
		TotalRows       int                      `json:"totalRows"`
		Items           []map[string]interface{} `json:"items"`
		NextPageToken   string                   `json:"nextPageToken"`
		HeuristicSearch bool                     `json:"heuristicSearch"`
	} `json:"data"`
}
