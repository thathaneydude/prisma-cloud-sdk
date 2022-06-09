package cwpp

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"net/url"
	"time"
)

const hostsEndpoint = "/hosts"

func (c *CwppClient) ListHosts(query HostsQuery) ([]Host, error) {
	var encoder = schema.NewEncoder()
	params := url.Values{}
	err := encoder.Encode(query, params)
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to decode Host query provided: %v", err)}
	}
	var hostsResponse []Host
	err = c.getWithResponseInterface(hostsEndpoint, params, &hostsResponse)
	if err != nil {
		return nil, err
	}
	return hostsResponse, nil
}

type HostsQuery struct {
	Offset        string   `schema:"offset,omitempty"`
	Limit         string   `schema:"limit,omitempty"`
	Search        string   `schema:"search,omitempty"`
	Sort          string   `schema:"sort,omitempty"`
	Reverse       bool     `schema:"reverse,omitempty"`
	Collections   []string `schema:"collections,omitempty"`
	AccountIDs    []string `schema:"accountIDs,omitempty"`
	Fields        []string `schema:"fields,omitempty"`
	Hostname      []string `schema:"hostname,omitempty"`
	Distro        []string `schema:"distro,omitempty"`
	Provider      []string `schema:"provider,omitempty"`
	Compact       bool     `schema:"compact,omitempty"`
	Clusters      []string `schema:"clusters,omitempty"`
	ComplianceIds []int    `schema:"complianceIds,omitempty"`
	Agentless     bool     `schema:"agentless,omitempty"`
	Stopped       bool     `schema:"stopped,empty"`
}

type Host struct {
	Id       string    `json:"_id"`
	Type     string    `json:"type"`
	Hostname string    `json:"hostname"`
	ScanTime time.Time `json:"scanTime"`
	Binaries []struct {
		Name     string   `json:"name"`
		Path     string   `json:"path"`
		Md5      string   `json:"md5"`
		CveCount int      `json:"cveCount"`
		Services []string `json:"services,omitempty"`
	} `json:"binaries"`
	Secrets         interface{}   `json:"Secrets"`
	StartupBinaries []interface{} `json:"startupBinaries"`
	OsDistro        string        `json:"osDistro"`
	OsDistroVersion string        `json:"osDistroVersion"`
	OsDistroRelease string        `json:"osDistroRelease"`
	Distro          string        `json:"distro"`
	Packages        []struct {
		PkgsType string `json:"pkgsType"`
		Pkgs     []struct {
			Version    string   `json:"version"`
			Name       string   `json:"name"`
			CveCount   int      `json:"cveCount"`
			License    string   `json:"license"`
			LayerTime  int      `json:"layerTime"`
			BinaryPkgs []string `json:"binaryPkgs,omitempty"`
			BinaryIdx  []int    `json:"binaryIdx,omitempty"`
		} `json:"pkgs"`
	} `json:"packages"`
	Files          interface{} `json:"files"`
	PackageManager bool        `json:"packageManager"`
	Applications   []struct {
		Name                 string `json:"name"`
		Version              string `json:"version"`
		Path                 string `json:"path"`
		LayerTime            int    `json:"layerTime"`
		KnownVulnerabilities int    `json:"knownVulnerabilities"`
	} `json:"applications"`
	Image struct {
		Created time.Time `json:"created"`
	} `json:"image"`
	History          []interface{} `json:"history"`
	ComplianceIssues interface{}   `json:"complianceIssues"`
	AllCompliance    struct {
		Compliance []struct {
			Text           string      `json:"text"`
			Id             int         `json:"id"`
			Severity       string      `json:"severity"`
			Cvss           int         `json:"cvss"`
			Status         string      `json:"status"`
			Cve            string      `json:"cve"`
			Cause          string      `json:"cause"`
			Description    string      `json:"description"`
			Title          string      `json:"title"`
			VecStr         string      `json:"vecStr"`
			Exploit        string      `json:"exploit"`
			RiskFactors    interface{} `json:"riskFactors"`
			Link           string      `json:"link"`
			Type           string      `json:"type"`
			PackageName    string      `json:"packageName"`
			PackageVersion string      `json:"packageVersion"`
			LayerTime      int         `json:"layerTime"`
			Templates      []string    `json:"templates"`
			Twistlock      bool        `json:"twistlock"`
			Cri            bool        `json:"cri"`
			Published      int         `json:"published"`
			FixDate        int         `json:"fixDate"`
			Discovered     time.Time   `json:"discovered"`
			FunctionLayer  string      `json:"functionLayer"`
			Custom         bool        `json:"custom,omitempty"`
		} `json:"compliance"`
		Enabled bool `json:"enabled"`
	} `json:"allCompliance"`
	Vulnerabilities []struct {
		Text        string  `json:"text"`
		Id          int     `json:"id"`
		Severity    string  `json:"severity"`
		Cvss        float64 `json:"cvss"`
		Status      string  `json:"status"`
		Cve         string  `json:"cve"`
		Cause       string  `json:"cause"`
		Description string  `json:"description"`
		Title       string  `json:"title"`
		VecStr      string  `json:"vecStr"`
		Exploit     string  `json:"exploit"`
		RiskFactors struct {
			HasFix struct {
			} `json:"Has fix,omitempty"`
			MediumSeverity struct {
			} `json:"Medium severity,omitempty"`
			PackageInUse struct {
			} `json:"Package in use,omitempty"`
			HighSeverity struct {
			} `json:"High severity,omitempty"`
			CriticalSeverity struct {
			} `json:"Critical severity,omitempty"`
			AttackComplexityLow struct {
			} `json:"Attack complexity: low,omitempty"`
			AttackVectorNetwork struct {
			} `json:"Attack vector: network,omitempty"`
			RecentVulnerability struct {
			} `json:"Recent vulnerability,omitempty"`
			DoS struct {
			} `json:"DoS,omitempty"`
		} `json:"riskFactors"`
		Link            string      `json:"link"`
		Type            string      `json:"type"`
		PackageName     string      `json:"packageName"`
		PackageVersion  string      `json:"packageVersion"`
		LayerTime       int         `json:"layerTime"`
		Templates       interface{} `json:"templates"`
		Twistlock       bool        `json:"twistlock"`
		Cri             bool        `json:"cri"`
		Published       int         `json:"published"`
		FixDate         int         `json:"fixDate"`
		ApplicableRules []string    `json:"applicableRules"`
		Discovered      time.Time   `json:"discovered"`
		BinaryPkgs      []string    `json:"binaryPkgs,omitempty"`
		FunctionLayer   string      `json:"functionLayer"`
	} `json:"vulnerabilities"`
	RepoTag                   interface{}   `json:"repoTag"`
	Tags                      []interface{} `json:"tags"`
	RepoDigests               []interface{} `json:"repoDigests"`
	CreationTime              time.Time     `json:"creationTime"`
	VulnerabilitiesCount      int           `json:"vulnerabilitiesCount"`
	ComplianceIssuesCount     int           `json:"complianceIssuesCount"`
	VulnerabilityDistribution struct {
		Critical int `json:"critical"`
		High     int `json:"high"`
		Medium   int `json:"medium"`
		Low      int `json:"low"`
		Total    int `json:"total"`
	} `json:"vulnerabilityDistribution"`
	ComplianceDistribution struct {
		Critical int `json:"critical"`
		High     int `json:"high"`
		Medium   int `json:"medium"`
		Low      int `json:"low"`
		Total    int `json:"total"`
	} `json:"complianceDistribution"`
	VulnerabilityRiskScore int    `json:"vulnerabilityRiskScore"`
	ComplianceRiskScore    int    `json:"complianceRiskScore"`
	EcsClusterName         string `json:"ecsClusterName"`
	RiskFactors            struct {
		AttackComplexityLow struct {
		} `json:"Attack complexity: low"`
		AttackVectorNetwork struct {
		} `json:"Attack vector: network"`
		CriticalSeverity struct {
		} `json:"Critical severity"`
		DoS struct {
		} `json:"DoS"`
		HasFix struct {
		} `json:"Has fix"`
		HighSeverity struct {
		} `json:"High severity"`
		MediumSeverity struct {
		} `json:"Medium severity"`
		PackageInUse struct {
		} `json:"Package in use"`
		RecentVulnerability struct {
		} `json:"Recent vulnerability"`
	} `json:"riskFactors"`
	Labels            []string `json:"labels"`
	InstalledProducts struct {
		Docker            string `json:"docker"`
		OsDistro          string `json:"osDistro"`
		HasPackageManager bool   `json:"hasPackageManager"`
	} `json:"installedProducts"`
	HostDevices []struct {
		Name string `json:"name"`
		Ip   string `json:"ip"`
	} `json:"hostDevices"`
	FirstScanTime time.Time `json:"firstScanTime"`
	CloudMetadata struct {
		ResourceID string `json:"resourceID"`
		Provider   string `json:"provider"`
		Type       string `json:"type"`
		Region     string `json:"region"`
		AccountID  string `json:"accountID"`
		Image      string `json:"image"`
	} `json:"cloudMetadata"`
	Clusters  []string      `json:"clusters"`
	Instances []interface{} `json:"instances"`
	Hosts     struct {
	} `json:"hosts"`
	Err                string   `json:"err"`
	Collections        []string `json:"collections"`
	ScanID             int      `json:"scanID"`
	TrustStatus        string   `json:"trustStatus"`
	FirewallProtection struct {
		Enabled   bool `json:"enabled"`
		Supported bool `json:"supported"`
	} `json:"firewallProtection"`
	AppEmbedded   bool        `json:"appEmbedded"`
	WildFireUsage interface{} `json:"wildFireUsage"`
	Agentless     bool        `json:"agentless"`
}
