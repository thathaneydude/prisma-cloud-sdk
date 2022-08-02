package cwpp

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/schema"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
)

const defendedEndpoint = "/defenders"

// ListDefenders retrieves all deployed Defenders with the required DefenderQuery.
//
// https://prisma.pan.dev/api/cloud/cwpp/defenders#operation/get-defenders
func (c *CwppClient) ListDefenders(query DefenderQuery) ([]Defender, error) {
	var encoder = schema.NewEncoder()
	params := url.Values{}
	err := encoder.Encode(query, params)
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to decode Defender query provided: %v", err)}
	}
	var defendersResponse []Defender
	err = c.getWithResponseInterface(defendedEndpoint, params, &defendersResponse)
	if err != nil {
		return nil, err
	}
	return defendersResponse, nil
}

type DefenderQuery struct {
	Offset        string   `json:"offset,omitempty"`
	Limit         string   `schema:"limit,omitempty"`
	Search        string   `schema:"search,omitempty"`
	Sort          string   `schema:"sort,omitempty"`
	Reverse       bool     `schema:"reverse,omitempty"`
	Collections   []string `schema:"collections,omitempty"`
	AccountIDs    []string `schema:"accountIDs,omitempty"`
	Fields        []string `schema:"fields,omitempty"`
	Hostname      []string `schema:"hostname,omitempty"`
	Role          string   `schema:"role,omitempty"`
	Connected     bool     `schema:"connected,omitempty"`
	Type          []string `schema:"type,omitempty"`
	Latest        bool     `schema:"latest,omitempty"`
	Cluster       []string `schema:"cluster,omitempty"`
	TasClusterIds []string `schema:"tasClusterIDs,omitempty"`
	UsingOldCA    bool     `schema:"usingOldCA,omitempty"`
}

type Defender struct {
	Category              string    `json:"category,omitempty"`
	CertificateExpiration time.Time `json:"certificateExpiration,omitempty"`
	CloudMetadata         struct {
		AccountID string `json:"accountID,omitempty"`
		Image     string `json:"image,omitempty"`
		Labels    []struct {
			Key        string    `json:"key,omitempty"`
			SourceName string    `json:"sourceName,omitempty"`
			SourceType string    `json:"sourceType,omitempty"`
			Timestamp  time.Time `json:"timestamp,omitempty"`
			Value      string    `json:"value,omitempty"`
		} `json:"labels,omitempty"`
		Name        string `json:"name,omitempty"`
		Provider    string `json:"provider,omitempty"`
		Region      string `json:"region,omitempty"`
		ResourceID  string `json:"resourceID,omitempty"`
		ResourceURL string `json:"resourceURL,omitempty"`
		Type        string `json:"type,omitempty"`
	} `json:"cloudMetadata,omitempty"`
	Cluster           string   `json:"cluster,omitempty"`
	ClusterID         string   `json:"clusterID,omitempty"`
	Collections       []string `json:"collections,omitempty"`
	CompatibleVersion bool     `json:"compatibleVersion,omitempty"`
	Connected         bool     `json:"connected,omitempty"`
	Features          struct {
		ClusterMonitoring bool   `json:"clusterMonitoring,omitempty"`
		ProxyListenerType string `json:"proxyListenerType,omitempty"`
	} `json:"features,omitempty"`
	FirewallProtection struct {
		Enabled              bool  `json:"enabled,omitempty"`
		Ports                []int `json:"ports,omitempty"`
		Supported            bool  `json:"supported,omitempty"`
		UnprotectedProcesses []struct {
			Port    int    `json:"port,omitempty"`
			Process string `json:"process,omitempty"`
		} `json:"unprotectedProcesses,omitempty"`
	} `json:"firewallProtection,omitempty"`
	Fqdn         string    `json:"fqdn,omitempty"`
	Hostname     string    `json:"hostname,omitempty"`
	LastModified time.Time `json:"lastModified,omitempty"`
	Port         int       `json:"port,omitempty"`
	Proxy        struct {
		Ca        string `json:"ca,omitempty"`
		HttpProxy string `json:"httpProxy,omitempty"`
		NoProxy   string `json:"noProxy,omitempty"`
		Password  struct {
			Encrypted string `json:"encrypted,omitempty"`
			Plain     string `json:"plain,omitempty"`
		} `json:"password,omitempty"`
		User string `json:"user,omitempty"`
	} `json:"proxy,omitempty"`
	RemoteLoggingSupported bool `json:"remoteLoggingSupported,omitempty"`
	RemoteMgmtSupported    bool `json:"remoteMgmtSupported,omitempty"`
	Status                 struct {
		AppFirewall struct {
			Enabled  bool   `json:"enabled,omitempty"`
			Err      string `json:"err,omitempty"`
			Hostname string `json:"hostname,omitempty"`
		} `json:"appFirewall,omitempty"`
		Container struct {
			Completed bool      `json:"completed,omitempty"`
			Errors    []string  `json:"errors,omitempty"`
			Hostname  string    `json:"hostname,omitempty"`
			ScanTime  time.Time `json:"scanTime,omitempty"`
			Scanning  bool      `json:"scanning,omitempty"`
			Selective bool      `json:"selective,omitempty"`
		} `json:"container,omitempty"`
		ContainerNetworkFirewall struct {
			Enabled  bool   `json:"enabled,omitempty"`
			Err      string `json:"err,omitempty"`
			Hostname string `json:"hostname,omitempty"`
		} `json:"containerNetworkFirewall,omitempty"`
		Features struct {
			Enabled  bool   `json:"enabled,omitempty"`
			Err      string `json:"err,omitempty"`
			Hostname string `json:"hostname,omitempty"`
		} `json:"features,omitempty"`
		Filesystem struct {
			Enabled  bool   `json:"enabled,omitempty"`
			Err      string `json:"err,omitempty"`
			Hostname string `json:"hostname,omitempty"`
		} `json:"filesystem,omitempty"`
		HostCustomCompliance struct {
			Enabled  bool   `json:"enabled,omitempty"`
			Err      string `json:"err,omitempty"`
			Hostname string `json:"hostname,omitempty"`
		} `json:"hostCustomCompliance,omitempty"`
		HostNetworkFirewall struct {
			Enabled  bool   `json:"enabled,omitempty"`
			Err      string `json:"err,omitempty"`
			Hostname string `json:"hostname,omitempty"`
		} `json:"hostNetworkFirewall,omitempty"`
		Image struct {
			Completed bool      `json:"completed,omitempty"`
			Errors    []string  `json:"errors,omitempty"`
			Hostname  string    `json:"hostname,omitempty"`
			ScanTime  time.Time `json:"scanTime,omitempty"`
			Scanning  bool      `json:"scanning,omitempty"`
			Selective bool      `json:"selective,omitempty"`
		} `json:"image,omitempty"`
		LastModified time.Time `json:"lastModified,omitempty"`
		Network      struct {
			Enabled  bool   `json:"enabled,omitempty"`
			Err      string `json:"err,omitempty"`
			Hostname string `json:"hostname,omitempty"`
		} `json:"network,omitempty"`
		Process struct {
			Enabled  bool   `json:"enabled,omitempty"`
			Err      string `json:"err,omitempty"`
			Hostname string `json:"hostname,omitempty"`
		} `json:"process,omitempty"`
		Runc struct {
			Enabled  bool   `json:"enabled,omitempty"`
			Err      string `json:"err,omitempty"`
			Hostname string `json:"hostname,omitempty"`
		} `json:"runc,omitempty"`
		Runtime struct {
			Enabled  bool   `json:"enabled,omitempty"`
			Err      string `json:"err,omitempty"`
			Hostname string `json:"hostname,omitempty"`
		} `json:"runtime,omitempty"`
		TasDroplets struct {
			Completed bool      `json:"completed,omitempty"`
			Errors    []string  `json:"errors,omitempty"`
			Hostname  string    `json:"hostname,omitempty"`
			ScanTime  time.Time `json:"scanTime,omitempty"`
			Scanning  bool      `json:"scanning,omitempty"`
			Selective bool      `json:"selective,omitempty"`
		} `json:"tasDroplets,omitempty"`
		Upgrade struct {
			Err          string    `json:"err,omitempty"`
			Hostname     string    `json:"hostname,omitempty"`
			LastModified time.Time `json:"lastModified,omitempty"`
			Progress     int       `json:"progress,omitempty"`
		} `json:"upgrade,omitempty"`
	} `json:"status,omitempty"`
	SystemInfo struct {
		CpuCount         int     `json:"cpuCount,omitempty"`
		FreeDiskSpaceGB  float32 `json:"freeDiskSpaceGB,omitempty"`
		KernelVersion    string  `json:"kernelVersion,omitempty"`
		MemoryGB         float32 `json:"memoryGB,omitempty"`
		TotalDiskSpaceGB float32 `json:"totalDiskSpaceGB,omitempty"`
	} `json:"systemInfo,omitempty"`
	TasClusterID string `json:"tasClusterID,omitempty"`
	Type         string `json:"type,omitempty"`
	UsingOldCA   bool   `json:"usingOldCA,omitempty"`
	Version      string `json:"version,omitempty"`
}
