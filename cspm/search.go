package cspm

import "github.com/thathaneydude/prisma-cloud-sdk/internal"

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
	Id                string `json:"id"`
	Limit             int    `json:"limit"`
	Query             string `json:"query"`
	SearchDescription string `json:"searchDescription"`
	SearchName        string `json:"searchName"`
	Sort              []struct {
		Direction string `json:"direction"`
		Field     string `json:"field"`
	} `json:"sort"`
	TimeRange struct {
		RelativeTimeType string `json:"relativeTimeType"`
		Type             string `json:"type"`
		Value            struct {
			Amount int    `json:"amount"`
			Unit   string `json:"unit"`
		} `json:"value"`
	} `json:"timeRange"`
	WithResourceJson bool `json:"withResourceJson"`
}

type SearchConfigResponse struct {
	AlertId        string                 `json:"alertId"`
	Async          bool                   `json:"async"`
	AsyncResultUrl string                 `json:"asyncResultUrl"`
	CloudType      string                 `json:"cloudType"`
	Cursor         int                    `json:"cursor"`
	Data           map[string]interface{} `json:"data"`
	Default        bool                   `json:"default"`
	Description    string                 `json:"description"`
	Filters        []struct {
		Name     string `json:"name"`
		Operator string `json:"operator"`
		Value    string `json:"value"`
	} `json:"filters"`
	GroupBy         []string `json:"groupBy"`
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	Query           string   `json:"query"`
	Saved           bool     `json:"saved"`
	SearchType      string   `json:"searchType"`
	TimeGranularity string   `json:"timeGranularity"`
	TimeRange       struct {
		RelativeTimeType string `json:"relativeTimeType"`
		Type             string `json:"type"`
		Value            struct {
			Amount int    `json:"amount"`
			Unit   string `json:"unit"`
		} `json:"value"`
	} `json:"timeRange"`
}
