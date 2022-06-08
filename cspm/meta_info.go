package cspm

const metaInfoEndpoint = "/meta_info"

// GetMetaInfo fetches the metadata pertaining to the authenticated JWT. Used primarily to determine the Twistlock URL
// for CWPP SaaS
func (c *CspmClient) GetMetaInfo() (*MetaInfoResponse, error) {
	var metaInfoResponse MetaInfoResponse
	err := c.getWithResponseInterface(metaInfoEndpoint, nil, &metaInfoResponse)
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
