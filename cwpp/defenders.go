package cwpp

import "time"

func (c *CwppClient) ListDefenders() {

}

type DefendersList struct {
	Category              []string  `json:"category"`
	CertificateExpiration time.Time `json:"certificateExpiration"`
	CloudMetadata         struct {
		AccountID string `json:"accountID"`
		Image     string `json:"image"`
		Labels    []struct {
			Key        string    `json:"key"`
			SourceName string    `json:"sourceName"`
			SourceType []string  `json:"sourceType"`
			Timestamp  time.Time `json:"timestamp"`
			Value      string    `json:"value"`
		} `json:"labels"`
		Name        string   `json:"name"`
		Provider    []string `json:"provider"`
		Region      string   `json:"region"`
		ResourceID  string   `json:"resourceID"`
		ResourceURL string   `json:"resourceURL"`
		Type        string   `json:"type"`
	} `json:"cloudMetadata"`
	Cluster           string   `json:"cluster"`
	ClusterID         string   `json:"clusterID"`
	Collections       []string `json:"collections"`
	CompatibleVersion bool     `json:"compatibleVersion"`
	Connected         bool     `json:"connected"`
	Features          struct {
		ClusterMonitoring bool     `json:"clusterMonitoring"`
		ProxyListenerType []string `json:"proxyListenerType"`
	} `json:"features"`
	FirewallProtection struct {
		Enabled              bool  `json:"enabled"`
		Ports                []int `json:"ports"`
		Supported            bool  `json:"supported"`
		UnprotectedProcesses []struct {
			Port    int    `json:"port"`
			Process string `json:"process"`
		} `json:"unprotectedProcesses"`
	} `json:"firewallProtection"`
	Fqdn         string    `json:"fqdn"`
	Hostname     string    `json:"hostname"`
	LastModified time.Time `json:"lastModified"`
	Port         int       `json:"port"`
	Proxy        struct {
		Ca        string `json:"ca"`
		HttpProxy string `json:"httpProxy"`
		NoProxy   string `json:"noProxy"`
		Password  struct {
			Encrypted string `json:"encrypted"`
			Plain     string `json:"plain"`
		} `json:"password"`
		User string `json:"user"`
	} `json:"proxy"`
	RemoteLoggingSupported bool `json:"remoteLoggingSupported"`
	RemoteMgmtSupported    bool `json:"remoteMgmtSupported"`
	Status                 struct {
		AppFirewall struct {
			Enabled  bool   `json:"enabled"`
			Err      string `json:"err"`
			Hostname string `json:"hostname"`
		} `json:"appFirewall"`
		Container struct {
			Completed bool      `json:"completed"`
			Errors    []string  `json:"errors"`
			Hostname  string    `json:"hostname"`
			ScanTime  time.Time `json:"scanTime"`
			Scanning  bool      `json:"scanning"`
			Selective bool      `json:"selective"`
		} `json:"container"`
		ContainerNetworkFirewall struct {
			Enabled  bool   `json:"enabled"`
			Err      string `json:"err"`
			Hostname string `json:"hostname"`
		} `json:"containerNetworkFirewall"`
		Features struct {
			Enabled  bool   `json:"enabled"`
			Err      string `json:"err"`
			Hostname string `json:"hostname"`
		} `json:"features"`
		Filesystem struct {
			Enabled  bool   `json:"enabled"`
			Err      string `json:"err"`
			Hostname string `json:"hostname"`
		} `json:"filesystem"`
		HostCustomCompliance struct {
			Enabled  bool   `json:"enabled"`
			Err      string `json:"err"`
			Hostname string `json:"hostname"`
		} `json:"hostCustomCompliance"`
		HostNetworkFirewall struct {
			Enabled  bool   `json:"enabled"`
			Err      string `json:"err"`
			Hostname string `json:"hostname"`
		} `json:"hostNetworkFirewall"`
		Image struct {
			Completed bool      `json:"completed"`
			Errors    []string  `json:"errors"`
			Hostname  string    `json:"hostname"`
			ScanTime  time.Time `json:"scanTime"`
			Scanning  bool      `json:"scanning"`
			Selective bool      `json:"selective"`
		} `json:"image"`
		LastModified time.Time `json:"lastModified"`
		Network      struct {
			Enabled  bool   `json:"enabled"`
			Err      string `json:"err"`
			Hostname string `json:"hostname"`
		} `json:"network"`
		Process struct {
			Enabled  bool   `json:"enabled"`
			Err      string `json:"err"`
			Hostname string `json:"hostname"`
		} `json:"process"`
		Runc struct {
			Enabled  bool   `json:"enabled"`
			Err      string `json:"err"`
			Hostname string `json:"hostname"`
		} `json:"runc"`
		Runtime struct {
			Enabled  bool   `json:"enabled"`
			Err      string `json:"err"`
			Hostname string `json:"hostname"`
		} `json:"runtime"`
		TasDroplets struct {
			Completed bool      `json:"completed"`
			Errors    []string  `json:"errors"`
			Hostname  string    `json:"hostname"`
			ScanTime  time.Time `json:"scanTime"`
			Scanning  bool      `json:"scanning"`
			Selective bool      `json:"selective"`
		} `json:"tasDroplets"`
		Upgrade struct {
			Err          string    `json:"err"`
			Hostname     string    `json:"hostname"`
			LastModified time.Time `json:"lastModified"`
			Progress     int       `json:"progress"`
		} `json:"upgrade"`
	} `json:"status"`
	SystemInfo struct {
		CpuCount         int    `json:"cpuCount"`
		FreeDiskSpaceGB  int    `json:"freeDiskSpaceGB"`
		KernelVersion    string `json:"kernelVersion"`
		MemoryGB         int    `json:"memoryGB"`
		TotalDiskSpaceGB int    `json:"totalDiskSpaceGB"`
	} `json:"systemInfo"`
	TasClusterID string   `json:"tasClusterID"`
	Type         []string `json:"type"`
	UsingOldCA   bool     `json:"usingOldCA"`
	Version      string   `json:"version"`
}
