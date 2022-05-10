package cspm

const metaInfoEndpoint = "/meta_info"

func (c *CspmClient) GetMetaInfo() (*MetaInfoResponse, error) {
	var metaInfoResponse MetaInfoResponse
	err := c.GetWithResponseInterface(metaInfoEndpoint, nil, &metaInfoResponse)
	if err != nil {
		return nil, err
	}

	return &metaInfoResponse, nil
}

type MetaInfoResponse struct {
	LicenseType  string `json:"licenseType,omitempty"`
	Marketplace  string `json:"marketplace,omitempty"`
	StartTs      int64  `json:"startTs,omitempty"`
	EndTs        int64  `json:"endTs,omitempty"`
	TwistlockUrl string `json:"twistlockUrl,omitempty"`
}
