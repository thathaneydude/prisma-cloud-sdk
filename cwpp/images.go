package cwpp

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"net/url"
	"time"
)

const imagesEndpoint = "/images"

// ListImages retrieves image scan reports
//
// Request requires an ImageQuery. By default, the resource fetches the first 50 images. Pagination can be done using
// the Offset & Limit attributes of the ImageQuery.
//
// https://prisma.pan.dev/api/cloud/cwpp/images#operation/get-images
func (c *CwppClient) ListImages(query ImageQuery) ([]Image, error) {
	var encoder = schema.NewEncoder()
	params := url.Values{}
	err := encoder.Encode(query, params)
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to decode query provided: %v", err)}
	}
	var imagesResponse []Image
	err = c.getWithResponseInterface(imagesEndpoint, params, &imagesResponse)
	if err != nil {
		return nil, err
	}
	return imagesResponse, nil
}

type ImageQuery struct {
	Offset          string   `schema:"offset,omitempty"`
	Limit           string   `schema:"limit,omitempty"`
	Search          string   `schema:"search,omitempty"`
	Sort            string   `schema:"sort,omitempty"`
	Reverse         bool     `schema:"reverse,omitempty"`
	Collections     []string `schema:"collections,omitempty"`
	AccountIDs      []string `schema:"accountIDs,omitempty"`
	Fields          []string `schema:"fields,omitempty"`
	Id              []string `schema:"id,omitempty"`
	Hostname        []string `schema:"hostname,omitempty"`
	Repository      []string `schema:"repository,omitempty"`
	Registry        []string `schema:"registry,omitempty"`
	Name            []string `schema:"name,omitempty"`
	Layers          bool     `schema:"layers,omitempty"`
	FilterBaseImage bool     `schema:"filterBaseImage,omitempty"`
	Compact         bool     `schema:"compact,omitempty"`
	TrustStatuses   []string `schema:"trustStatuses,omitempty"`
	Clusters        []string `schema:"clusters,omitempty"`
	ComplianceIds   []int    `schema:"complianceIds,omitempty"`
	AppEmbedded     bool     `schema:"appEmbedded,omitempty"`
}

type Image struct {
	Id       string    `json:"_id,omitempty"`
	Type     string    `json:"type,omitempty"`
	Hostname string    `json:"hostname,omitempty"`
	ScanTime time.Time `json:"scanTime,omitempty"`
	Binaries []struct {
		Name       string `json:"name,omitempty"`
		Path       string `json:"path,omitempty"`
		Md5        string `json:"md5,omitempty"`
		CveCount   int    `json:"cveCount,omitempty"`
		Version    string `json:"version,omitempty"`
		MissingPkg bool   `json:"missingPkg,omitempty"`
	} `json:"binaries,omitempty"`
	Secrets         []string `json:"Secrets,omitempty"`
	StartupBinaries []struct {
		Name     string `json:"name,omitempty"`
		Path     string `json:"path,omitempty"`
		Md5      string `json:"md5,omitempty"`
		CveCount int    `json:"cveCount,omitempty"`
	} `json:"startupBinaries,omitempty"`
	OsDistro        string `json:"osDistro,omitempty"`
	OsDistroVersion string `json:"osDistroVersion,omitempty"`
	OsDistroRelease string `json:"osDistroRelease,omitempty"`
	Distro          string `json:"distro,omitempty"`
	Packages        []struct {
		PkgsType string `json:"pkgsType,omitempty"`
		Pkgs     []struct {
			Version    string   `json:"version,omitempty"`
			Name       string   `json:"name,omitempty"`
			CveCount   int      `json:"cveCount,omitempty"`
			License    string   `json:"license,omitempty"`
			LayerTime  int      `json:"layerTime,omitempty"`
			BinaryPkgs []string `json:"binaryPkgs,omitempty"`
		} `json:"pkgs,omitempty"`
	} `json:"packages,omitempty"`
	Files          []interface{} `json:"files,omitempty"`
	PackageManager bool          `json:"packageManager,omitempty"`
	Applications   []struct {
		Name                 string `json:"name,omitempty"`
		Version              string `json:"version,omitempty"`
		Path                 string `json:"path,omitempty"`
		LayerTime            int    `json:"layerTime,omitempty"`
		KnownVulnerabilities int    `json:"knownVulnerabilities,omitempty"`
	} `json:"applications,omitempty"`
	Image struct {
		Created time.Time `json:"created,omitempty"`
	} `json:"image,omitempty"`
	History []struct {
		Created     int      `json:"created,omitempty"`
		Instruction string   `json:"instruction,omitempty"`
		SizeBytes   int      `json:"sizeBytes,omitempty"`
		Id          string   `json:"id,omitempty"`
		EmptyLayer  bool     `json:"emptyLayer,omitempty"`
		Tags        []string `json:"tags,omitempty,omitempty"`
	} `json:"history,omitempty"`
	Id1              string `json:"id,omitempty"`
	ComplianceIssues []struct {
		Text           string      `json:"text,omitempty"`
		Id             int         `json:"id,omitempty"`
		Severity       string      `json:"severity,omitempty"`
		Cvss           int         `json:"cvss,omitempty"`
		Status         string      `json:"status,omitempty"`
		Cve            string      `json:"cve,omitempty"`
		Cause          string      `json:"cause,omitempty"`
		Description    string      `json:"description,omitempty"`
		Title          string      `json:"title,omitempty"`
		VecStr         string      `json:"vecStr,omitempty"`
		Exploit        string      `json:"exploit,omitempty"`
		RiskFactors    interface{} `json:"riskFactors,omitempty"`
		Link           string      `json:"link,omitempty"`
		Type           string      `json:"type,omitempty"`
		PackageName    string      `json:"packageName,omitempty"`
		PackageVersion string      `json:"packageVersion,omitempty"`
		LayerTime      int         `json:"layerTime,omitempty"`
		Templates      []string    `json:"templates,omitempty"`
		Twistlock      bool        `json:"twistlock,omitempty"`
		Cri            bool        `json:"cri,omitempty"`
		Published      int         `json:"published,omitempty"`
		FixDate        int         `json:"fixDate,omitempty"`
		Discovered     time.Time   `json:"discovered,omitempty"`
		FunctionLayer  string      `json:"functionLayer,omitempty"`
	} `json:"complianceIssues,omitempty"`
	AllCompliance struct {
	} `json:"allCompliance,omitempty"`
	Vulnerabilities []struct {
		Text        string  `json:"text,omitempty"`
		Id          int     `json:"id,omitempty"`
		Severity    string  `json:"severity,omitempty"`
		Cvss        float64 `json:"cvss,omitempty"`
		Status      string  `json:"status,omitempty"`
		Cve         string  `json:"cve,omitempty"`
		Cause       string  `json:"cause,omitempty"`
		Description string  `json:"description,omitempty"`
		Title       string  `json:"title,omitempty"`
		VecStr      string  `json:"vecStr,omitempty"`
		Exploit     string  `json:"exploit,omitempty"`
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
		} `json:"riskFactors,omitempty"`
		Link            string      `json:"link,omitempty"`
		Type            string      `json:"type,omitempty"`
		PackageName     string      `json:"packageName,omitempty"`
		PackageVersion  string      `json:"packageVersion,omitempty"`
		LayerTime       int         `json:"layerTime,omitempty"`
		Templates       interface{} `json:"templates,omitempty"`
		Twistlock       bool        `json:"twistlock,omitempty"`
		Cri             bool        `json:"cri,omitempty"`
		Published       int         `json:"published,omitempty"`
		FixDate         int         `json:"fixDate,omitempty"`
		ApplicableRules []string    `json:"applicableRules,omitempty,omitempty"`
		Discovered      time.Time   `json:"discovered,omitempty"`
		FunctionLayer   string      `json:"functionLayer,omitempty"`
		VulnTagInfos    []struct {
			Name    string `json:"name,omitempty"`
			Color   string `json:"color,omitempty"`
			Comment string `json:"comment,omitempty"`
		} `json:"vulnTagInfos,omitempty"`
		BinaryPkgs []string `json:"binaryPkgs,omitempty"`
	} `json:"vulnerabilities,omitempty"`
	RepoTag struct {
		Registry string `json:"registry,omitempty"`
		Repo     string `json:"repo,omitempty"`
		Tag      string `json:"tag,omitempty"`
	} `json:"repoTag,omitempty"`
	Tags []struct {
		Registry string `json:"registry,omitempty"`
		Repo     string `json:"repo,omitempty"`
		Tag      string `json:"tag,omitempty"`
	} `json:"tags,omitempty"`
	RepoDigests               []string  `json:"repoDigests,omitempty"`
	CreationTime              time.Time `json:"creationTime,omitempty"`
	VulnerabilitiesCount      int       `json:"vulnerabilitiesCount,omitempty"`
	ComplianceIssuesCount     int       `json:"complianceIssuesCount,omitempty"`
	VulnerabilityDistribution struct {
		Critical int `json:"critical,omitempty"`
		High     int `json:"high,omitempty"`
		Medium   int `json:"medium,omitempty"`
		Low      int `json:"low,omitempty"`
		Total    int `json:"total,omitempty"`
	} `json:"vulnerabilityDistribution,omitempty"`
	ComplianceDistribution struct {
		Critical int `json:"critical,omitempty"`
		High     int `json:"high,omitempty"`
		Medium   int `json:"medium,omitempty"`
		Low      int `json:"low,omitempty"`
		Total    int `json:"total,omitempty"`
	} `json:"complianceDistribution,omitempty"`
	VulnerabilityRiskScore int      `json:"vulnerabilityRiskScore,omitempty"`
	ComplianceRiskScore    int      `json:"complianceRiskScore,omitempty"`
	Layers                 []string `json:"layers,omitempty"`
	TopLayer               string   `json:"topLayer,omitempty"`
	RiskFactors            struct {
		AttackComplexityLow struct {
		} `json:"Attack complexity: low,omitempty"`
		AttackVectorNetwork struct {
		} `json:"Attack vector: network,omitempty"`
		CriticalSeverity struct {
		} `json:"Critical severity,omitempty"`
		DoS struct {
		} `json:"DoS,omitempty"`
		ExploitExists struct {
		} `json:"Exploit exists,omitempty"`
		HasFix struct {
		} `json:"Has fix,omitempty"`
		HighSeverity struct {
		} `json:"High severity,omitempty"`
		MediumSeverity struct {
		} `json:"Medium severity,omitempty"`
		RecentVulnerability struct {
		} `json:"Recent vulnerability,omitempty"`
		RemoteExecution struct {
		} `json:"Remote execution,omitempty"`
	} `json:"riskFactors,omitempty"`
	Labels            []string `json:"labels,omitempty"`
	InstalledProducts struct {
		Docker            string `json:"docker,omitempty"`
		OsDistro          string `json:"osDistro,omitempty"`
		HasPackageManager bool   `json:"hasPackageManager,omitempty"`
	} `json:"installedProducts,omitempty"`
	ScanVersion   string    `json:"scanVersion,omitempty"`
	FirstScanTime time.Time `json:"firstScanTime,omitempty"`
	CloudMetadata struct {
		AccountID string `json:"accountID,omitempty"`
	} `json:"cloudMetadata,omitempty"`
	Namespaces     []string `json:"namespaces,omitempty"`
	ExternalLabels []struct {
		SourceType string    `json:"sourceType,omitempty"`
		SourceName string    `json:"sourceName,omitempty"`
		Timestamp  time.Time `json:"timestamp,omitempty"`
		Key        string    `json:"key,omitempty"`
		Value      string    `json:"value,omitempty"`
	} `json:"externalLabels,omitempty"`
	Clusters  []string `json:"clusters,omitempty"`
	Instances []struct {
		Image    string    `json:"image,omitempty"`
		Host     string    `json:"host,omitempty"`
		Modified time.Time `json:"modified,omitempty"`
		Tag      string    `json:"tag,omitempty"`
		Repo     string    `json:"repo,omitempty"`
		Registry string    `json:"registry,omitempty"`
	} `json:"instances,omitempty"`
	Hosts              map[string]interface{} `json:"hosts,omitempty"`
	Err                string                 `json:"err,omitempty"`
	Collections        []string               `json:"collections,omitempty"`
	ScanID             int                    `json:"scanID,omitempty"`
	TrustStatus        string                 `json:"trustStatus,omitempty"`
	FirewallProtection struct {
		Enabled   bool `json:"enabled,omitempty"`
		Supported bool `json:"supported,omitempty"`
	} `json:"firewallProtection,omitempty"`
	AppEmbedded   bool        `json:"appEmbedded,omitempty"`
	WildFireUsage interface{} `json:"wildFireUsage,omitempty"`
	Agentless     bool        `json:"agentless,omitempty"`
}
