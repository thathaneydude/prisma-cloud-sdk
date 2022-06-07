package cs

import (
	"fmt"
	"net/url"
	"time"
)

const listRepositoriesEndpoint = "/code/api/v1/repositories"

// ListRepositories Returns a list of repositories that are integrated with Prisma Cloud Code Security.
// errorCounts:  true = Include error counts
// https://prisma.pan.dev/api/cloud/code/code-security#operation/getRepositories
func (c *CsClient) ListRepositories(errorsCount bool) ([]Repository, error) {
	var repos []Repository
	params := url.Values{}
	params.Set("errorCounts", fmt.Sprintf("%v", errorsCount))

	err := c.getWithResponseInterface(listRepositoriesEndpoint, params, &repos)
	if err != nil {
		return nil, err
	}
	return repos, nil
}

type Repository struct {
	Id            string    `json:"id,omitempty"`
	Repository    string    `json:"repository,omitempty"`
	Source        string    `json:"source,omitempty"`
	Owner         string    `json:"owner,omitempty"`
	DefaultBranch string    `json:"defaultBranch,omitempty"`
	IsPublic      bool      `json:"isPublic,omitempty"`
	Runs          int       `json:"runs,omitempty"`
	CreationDate  time.Time `json:"creationDate,omitempty"`
	LastScanDate  string    `json:"lastScanDate,omitempty"`
	Errors        int       `json:"errors,omitempty"`
}
