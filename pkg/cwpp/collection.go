package cwpp

const collectionEndpoint = "collections"

func (c *CwppClient) ListCollections() ([]Collection, error) {
	var collections []Collection
	err := c.GetWithResponseInterface(collectionEndpoint, nil, &collections)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

type Collection struct {
	AccountIds  []string `json:"accountIDs,omitempty"`
	AppIds      []string `json:"appIDs,omitempty"`
	Clusters    []string `json:"clusters,omitempty"`
	CodeRepos   []string `json:"codeRepos,omitempty"`
	Color       string   `json:"color,omitempty"`
	Containers  []string `json:"containers,omitempty"`
	Description string   `json:"description,omitempty"`
	Functions   []string `json:"functions,omitempty"`
	Hosts       []string `json:"hosts,omitempty"`
	Images      []string `json:"images,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Name        string   `json:"name,omitempty"`
	Namespaces  []string `json:"namespaces,omitempty"`
}
