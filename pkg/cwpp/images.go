package cwpp

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	"net/url"
	"prisma-cloud-sdk/pkg"
	"time"
)

const imagesEndpoint = "/images"

func (c *CwppClient) ListImages(query ImageQuery) ([]Image, error) {
	var encoder = schema.NewEncoder()
	params := url.Values{}
	err := encoder.Encode(query, params)
	if err != nil {
		return nil, &pkg.GenericError{Msg: fmt.Sprintf("Failed to decode query provided: %v", err)}
	}
	var imagesResponse []Image
	logrus.Infof("Encoded struct: %v", params)
	err = c.GetWithResponseInterface(imagesEndpoint, params, &imagesResponse)
	if err != nil {
		return nil, err
	}
	return imagesResponse, nil
}

type ImageQuery struct {
	Offset          string   `schema:"offset"`
	Limit           string   `schema:"limit"`
	Search          string   `schema:"search"`
	Sort            string   `schema:"sort"`
	Reverse         bool     `schema:"reverse"`
	Collections     []string `schema:"collections"`
	AccountIDs      []string `schema:"accountIDs"`
	Fields          []string `schema:"fields"`
	Id              []string `schema:"id"`
	Hostname        []string `schema:"hostname"`
	Repository      []string `schema:"repository"`
	Registry        []string `schema:"registry"`
	Name            []string `schema:"name"`
	Layers          bool     `schema:"layers"`
	FilterBaseImage bool     `schema:"filterBaseImage"`
	Compact         bool     `schema:"compact"`
	TrustStatuses   []string `schema:"trustStatuses"`
	Clusters        []string `schema:"clusters"`
	ComplianceIds   []int    `schema:"complianceIds"`
	AppEmbedded     bool     `schema:"appEmbedded"`
}

type Image struct {
	Id       string    `json:"_id"`
	Type     string    `json:"type"`
	Hostname string    `json:"hostname"`
	ScanTime time.Time `json:"scanTime"`
	Binaries []struct {
		Name       string `json:"name"`
		Path       string `json:"path"`
		Md5        string `json:"md5"`
		CveCount   int    `json:"cveCount"`
		Version    string `json:"version,omitempty"`
		MissingPkg bool   `json:"missingPkg,omitempty"`
	} `json:"binaries"`
	Secrets         []string `json:"Secrets"`
	StartupBinaries []struct {
		Name     string `json:"name"`
		Path     string `json:"path"`
		Md5      string `json:"md5"`
		CveCount int    `json:"cveCount"`
	} `json:"startupBinaries"`
	OsDistro        string `json:"osDistro"`
	OsDistroVersion string `json:"osDistroVersion"`
	OsDistroRelease string `json:"osDistroRelease"`
	Distro          string `json:"distro"`
	Packages        []struct {
		PkgsType string `json:"pkgsType"`
		Pkgs     []struct {
			Version    string   `json:"version"`
			Name       string   `json:"name"`
			CveCount   int      `json:"cveCount"`
			License    string   `json:"license"`
			LayerTime  int      `json:"layerTime"`
			BinaryPkgs []string `json:"binaryPkgs,omitempty"`
		} `json:"pkgs"`
	} `json:"packages"`
	Files          []interface{} `json:"files"`
	PackageManager bool          `json:"packageManager"`
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
	History []struct {
		Created     int      `json:"created"`
		Instruction string   `json:"instruction"`
		SizeBytes   int      `json:"sizeBytes,omitempty"`
		Id          string   `json:"id"`
		EmptyLayer  bool     `json:"emptyLayer"`
		Tags        []string `json:"tags,omitempty"`
	} `json:"history"`
	Id1              string `json:"id"`
	ComplianceIssues []struct {
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
	} `json:"complianceIssues"`
	AllCompliance struct {
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
			AttackComplexityLow struct {
			} `json:"Attack complexity: low,omitempty"`
			DoS struct {
			} `json:"DoS,omitempty"`
			AttackVectorNetwork struct {
			} `json:"Attack vector: network,omitempty"`
			HasFix struct {
			} `json:"Has fix,omitempty"`
			MediumSeverity struct {
			} `json:"Medium severity,omitempty"`
			RecentVulnerability struct {
			} `json:"Recent vulnerability,omitempty"`
			ExploitExists struct {
			} `json:"Exploit exists,omitempty"`
			HighSeverity struct {
			} `json:"High severity,omitempty"`
			RemoteExecution struct {
			} `json:"Remote execution,omitempty"`
			CriticalSeverity struct {
			} `json:"Critical severity,omitempty"`
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
		ApplicableRules []string    `json:"applicableRules,omitempty"`
		Discovered      time.Time   `json:"discovered"`
		FunctionLayer   string      `json:"functionLayer"`
		VulnTagInfos    []struct {
			Name    string `json:"name"`
			Color   string `json:"color"`
			Comment string `json:"comment,omitempty"`
		} `json:"vulnTagInfos,omitempty"`
		BinaryPkgs []string `json:"binaryPkgs,omitempty"`
	} `json:"vulnerabilities"`
	RepoTag struct {
		Registry string `json:"registry"`
		Repo     string `json:"repo"`
		Tag      string `json:"tag"`
	} `json:"repoTag"`
	Tags []struct {
		Registry string `json:"registry"`
		Repo     string `json:"repo"`
		Tag      string `json:"tag"`
	} `json:"tags"`
	RepoDigests               []string  `json:"repoDigests"`
	CreationTime              time.Time `json:"creationTime"`
	VulnerabilitiesCount      int       `json:"vulnerabilitiesCount"`
	ComplianceIssuesCount     int       `json:"complianceIssuesCount"`
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
	VulnerabilityRiskScore int      `json:"vulnerabilityRiskScore"`
	ComplianceRiskScore    int      `json:"complianceRiskScore"`
	Layers                 []string `json:"layers"`
	TopLayer               string   `json:"topLayer"`
	RiskFactors            struct {
		AttackComplexityLow struct {
		} `json:"Attack complexity: low"`
		AttackVectorNetwork struct {
		} `json:"Attack vector: network"`
		CriticalSeverity struct {
		} `json:"Critical severity"`
		DoS struct {
		} `json:"DoS"`
		ExploitExists struct {
		} `json:"Exploit exists"`
		HasFix struct {
		} `json:"Has fix"`
		HighSeverity struct {
		} `json:"High severity"`
		MediumSeverity struct {
		} `json:"Medium severity"`
		RecentVulnerability struct {
		} `json:"Recent vulnerability"`
		RemoteExecution struct {
		} `json:"Remote execution"`
	} `json:"riskFactors"`
	Labels            []string `json:"labels"`
	InstalledProducts struct {
		Docker            string `json:"docker"`
		OsDistro          string `json:"osDistro"`
		HasPackageManager bool   `json:"hasPackageManager"`
	} `json:"installedProducts"`
	ScanVersion   string    `json:"scanVersion"`
	FirstScanTime time.Time `json:"firstScanTime"`
	CloudMetadata struct {
		AccountID string `json:"accountID"`
	} `json:"cloudMetadata"`
	Namespaces     []string `json:"namespaces"`
	ExternalLabels []struct {
		SourceType string    `json:"sourceType"`
		SourceName string    `json:"sourceName"`
		Timestamp  time.Time `json:"timestamp"`
		Key        string    `json:"key"`
		Value      string    `json:"value"`
	} `json:"externalLabels"`
	Clusters  []string `json:"clusters"`
	Instances []struct {
		Image    string    `json:"image"`
		Host     string    `json:"host"`
		Modified time.Time `json:"modified"`
		Tag      string    `json:"tag"`
		Repo     string    `json:"repo"`
		Registry string    `json:"registry"`
	} `json:"instances"`
	Hosts struct {
		MasterSelfhostedPphonpaseuthDemoTwistlockCom struct {
			Modified   time.Time `json:"modified"`
			Cluster    string    `json:"cluster"`
			Namespaces []string  `json:"namespaces"`
			AccountID  string    `json:"accountID"`
		} `json:"master-selfhosted-pphonpaseuth-demo-twistlock-com"`
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
