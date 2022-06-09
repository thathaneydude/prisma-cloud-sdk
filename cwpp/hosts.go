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
	Secrets       []string `json:"Secrets"`
	Id            string   `json:"_id"`
	Agentless     bool     `json:"agentless"`
	AllCompliance struct {
		Compliance []struct {
			ApplicableRules []string  `json:"applicableRules"`
			BinaryPkgs      []string  `json:"binaryPkgs"`
			Block           bool      `json:"block"`
			Cause           string    `json:"cause"`
			Cri             bool      `json:"cri"`
			Custom          bool      `json:"custom"`
			Cve             string    `json:"cve"`
			Cvss            int       `json:"cvss"`
			Description     string    `json:"description"`
			Discovered      time.Time `json:"discovered"`
			Exploit         []string  `json:"exploit"`
			FixDate         int       `json:"fixDate"`
			FixLink         string    `json:"fixLink"`
			FunctionLayer   string    `json:"functionLayer"`
			GracePeriodDays int       `json:"gracePeriodDays"`
			Id              int       `json:"id"`
			LayerTime       int       `json:"layerTime"`
			Link            string    `json:"link"`
			PackageName     string    `json:"packageName"`
			PackageVersion  string    `json:"packageVersion"`
			Published       int       `json:"published"`
			RiskFactors     struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"riskFactors"`
			Severity     string     `json:"severity"`
			Status       string     `json:"status"`
			Templates    [][]string `json:"templates"`
			Text         string     `json:"text"`
			Title        string     `json:"title"`
			Twistlock    bool       `json:"twistlock"`
			Type         []string   `json:"type"`
			VecStr       string     `json:"vecStr"`
			VulnTagInfos []struct {
				Color   string `json:"color"`
				Comment string `json:"comment"`
				Name    string `json:"name"`
			} `json:"vulnTagInfos"`
		} `json:"compliance"`
		Enabled bool `json:"enabled"`
	} `json:"allCompliance"`
	AppEmbedded  bool `json:"appEmbedded"`
	Applications []struct {
		KnownVulnerabilities int    `json:"knownVulnerabilities"`
		LayerTime            int    `json:"layerTime"`
		Name                 string `json:"name"`
		Path                 string `json:"path"`
		Version              string `json:"version"`
	} `json:"applications"`
	BaseImage string `json:"baseImage"`
	Binaries  []struct {
		Altered       bool     `json:"altered"`
		CveCount      int      `json:"cveCount"`
		Deps          []string `json:"deps"`
		FunctionLayer string   `json:"functionLayer"`
		Md5           string   `json:"md5"`
		MissingPkg    bool     `json:"missingPkg"`
		Name          string   `json:"name"`
		Path          string   `json:"path"`
		PkgRootDir    string   `json:"pkgRootDir"`
		Services      []string `json:"services"`
		Version       string   `json:"version"`
	} `json:"binaries"`
	CloudMetadata struct {
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
		VmID        string   `json:"vmID"`
	} `json:"cloudMetadata"`
	Clusters               []string `json:"clusters"`
	Collections            []string `json:"collections"`
	ComplianceDistribution struct {
		Critical int `json:"critical"`
		High     int `json:"high"`
		Low      int `json:"low"`
		Medium   int `json:"medium"`
		Total    int `json:"total"`
	} `json:"complianceDistribution"`
	ComplianceIssues []struct {
		ApplicableRules []string  `json:"applicableRules"`
		BinaryPkgs      []string  `json:"binaryPkgs"`
		Block           bool      `json:"block"`
		Cause           string    `json:"cause"`
		Cri             bool      `json:"cri"`
		Custom          bool      `json:"custom"`
		Cve             string    `json:"cve"`
		Cvss            int       `json:"cvss"`
		Description     string    `json:"description"`
		Discovered      time.Time `json:"discovered"`
		Exploit         []string  `json:"exploit"`
		FixDate         int       `json:"fixDate"`
		FixLink         string    `json:"fixLink"`
		FunctionLayer   string    `json:"functionLayer"`
		GracePeriodDays int       `json:"gracePeriodDays"`
		Id              int       `json:"id"`
		LayerTime       int       `json:"layerTime"`
		Link            string    `json:"link"`
		PackageName     string    `json:"packageName"`
		PackageVersion  string    `json:"packageVersion"`
		Published       int       `json:"published"`
		RiskFactors     struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"riskFactors"`
		Severity     string     `json:"severity"`
		Status       string     `json:"status"`
		Templates    [][]string `json:"templates"`
		Text         string     `json:"text"`
		Title        string     `json:"title"`
		Twistlock    bool       `json:"twistlock"`
		Type         []string   `json:"type"`
		VecStr       string     `json:"vecStr"`
		VulnTagInfos []struct {
			Color   string `json:"color"`
			Comment string `json:"comment"`
			Name    string `json:"name"`
		} `json:"vulnTagInfos"`
	} `json:"complianceIssues"`
	ComplianceIssuesCount int       `json:"complianceIssuesCount"`
	ComplianceRiskScore   int       `json:"complianceRiskScore"`
	CreationTime          time.Time `json:"creationTime"`
	Distro                string    `json:"distro"`
	EcsClusterName        string    `json:"ecsClusterName"`
	Err                   string    `json:"err"`
	ExternalLabels        []struct {
		Key        string    `json:"key"`
		SourceName string    `json:"sourceName"`
		SourceType []string  `json:"sourceType"`
		Timestamp  time.Time `json:"timestamp"`
		Value      string    `json:"value"`
	} `json:"externalLabels"`
	Files []struct {
		Md5    string `json:"md5"`
		Path   string `json:"path"`
		Sha1   string `json:"sha1"`
		Sha256 string `json:"sha256"`
	} `json:"files"`
	FirewallProtection struct {
		Enabled              bool     `json:"enabled"`
		OutOfBandMode        []string `json:"outOfBandMode"`
		Ports                []int    `json:"ports"`
		Supported            bool     `json:"supported"`
		TlsPorts             []int    `json:"tlsPorts"`
		UnprotectedProcesses []struct {
			Port    int    `json:"port"`
			Process string `json:"process"`
			Tls     bool   `json:"tls"`
		} `json:"unprotectedProcesses"`
	} `json:"firewallProtection"`
	FirstScanTime time.Time `json:"firstScanTime"`
	History       []struct {
		BaseLayer       bool     `json:"baseLayer"`
		Created         int      `json:"created"`
		EmptyLayer      bool     `json:"emptyLayer"`
		Id              string   `json:"id"`
		Instruction     string   `json:"instruction"`
		SizeBytes       int      `json:"sizeBytes"`
		Tags            []string `json:"tags"`
		Vulnerabilities []struct {
			ApplicableRules []string  `json:"applicableRules"`
			BinaryPkgs      []string  `json:"binaryPkgs"`
			Block           bool      `json:"block"`
			Cause           string    `json:"cause"`
			Cri             bool      `json:"cri"`
			Custom          bool      `json:"custom"`
			Cve             string    `json:"cve"`
			Cvss            int       `json:"cvss"`
			Description     string    `json:"description"`
			Discovered      time.Time `json:"discovered"`
			Exploit         []string  `json:"exploit"`
			FixDate         int       `json:"fixDate"`
			FixLink         string    `json:"fixLink"`
			FunctionLayer   string    `json:"functionLayer"`
			GracePeriodDays int       `json:"gracePeriodDays"`
			Id              int       `json:"id"`
			LayerTime       int       `json:"layerTime"`
			Link            string    `json:"link"`
			PackageName     string    `json:"packageName"`
			PackageVersion  string    `json:"packageVersion"`
			Published       int       `json:"published"`
			RiskFactors     struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"riskFactors"`
			Severity     string     `json:"severity"`
			Status       string     `json:"status"`
			Templates    [][]string `json:"templates"`
			Text         string     `json:"text"`
			Title        string     `json:"title"`
			Twistlock    bool       `json:"twistlock"`
			Type         []string   `json:"type"`
			VecStr       string     `json:"vecStr"`
			VulnTagInfos []struct {
				Color   string `json:"color"`
				Comment string `json:"comment"`
				Name    string `json:"name"`
			} `json:"vulnTagInfos"`
		} `json:"vulnerabilities"`
	} `json:"history"`
	HostDevices []struct {
		Ip   string `json:"ip"`
		Name string `json:"name"`
	} `json:"hostDevices"`
	Hostname string `json:"hostname"`
	Hosts    struct {
		Property1 struct {
			AccountID   string    `json:"accountID"`
			AppEmbedded bool      `json:"appEmbedded"`
			Cluster     string    `json:"cluster"`
			Modified    time.Time `json:"modified"`
			Namespaces  []string  `json:"namespaces"`
		} `json:"property1"`
		Property2 struct {
			AccountID   string    `json:"accountID"`
			AppEmbedded bool      `json:"appEmbedded"`
			Cluster     string    `json:"cluster"`
			Modified    time.Time `json:"modified"`
			Namespaces  []string  `json:"namespaces"`
		} `json:"property2"`
	} `json:"hosts"`
	Id1   string `json:"id"`
	Image struct {
		Created     time.Time `json:"created"`
		Entrypoint  []string  `json:"entrypoint"`
		Env         []string  `json:"env"`
		Healthcheck bool      `json:"healthcheck"`
		History     []struct {
			BaseLayer       bool     `json:"baseLayer"`
			Created         int      `json:"created"`
			EmptyLayer      bool     `json:"emptyLayer"`
			Id              string   `json:"id"`
			Instruction     string   `json:"instruction"`
			SizeBytes       int      `json:"sizeBytes"`
			Tags            []string `json:"tags"`
			Vulnerabilities []struct {
				ApplicableRules []string  `json:"applicableRules"`
				BinaryPkgs      []string  `json:"binaryPkgs"`
				Block           bool      `json:"block"`
				Cause           string    `json:"cause"`
				Cri             bool      `json:"cri"`
				Custom          bool      `json:"custom"`
				Cve             string    `json:"cve"`
				Cvss            int       `json:"cvss"`
				Description     string    `json:"description"`
				Discovered      time.Time `json:"discovered"`
				Exploit         []string  `json:"exploit"`
				FixDate         int       `json:"fixDate"`
				FixLink         string    `json:"fixLink"`
				FunctionLayer   string    `json:"functionLayer"`
				GracePeriodDays int       `json:"gracePeriodDays"`
				Id              int       `json:"id"`
				LayerTime       int       `json:"layerTime"`
				Link            string    `json:"link"`
				PackageName     string    `json:"packageName"`
				PackageVersion  string    `json:"packageVersion"`
				Published       int       `json:"published"`
				RiskFactors     struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"riskFactors"`
				Severity     string     `json:"severity"`
				Status       string     `json:"status"`
				Templates    [][]string `json:"templates"`
				Text         string     `json:"text"`
				Title        string     `json:"title"`
				Twistlock    bool       `json:"twistlock"`
				Type         []string   `json:"type"`
				VecStr       string     `json:"vecStr"`
				VulnTagInfos []struct {
					Color   string `json:"color"`
					Comment string `json:"comment"`
					Name    string `json:"name"`
				} `json:"vulnTagInfos"`
			} `json:"vulnerabilities"`
		} `json:"history"`
		Id     string `json:"id"`
		Labels struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"labels"`
		Layers     []string `json:"layers"`
		Os         string   `json:"os"`
		RepoDigest []string `json:"repoDigest"`
		RepoTags   []string `json:"repoTags"`
		User       string   `json:"user"`
		WorkingDir string   `json:"workingDir"`
	} `json:"image"`
	InstalledProducts struct {
		Agentless                      bool   `json:"agentless"`
		Apache                         string `json:"apache"`
		AwsCloud                       bool   `json:"awsCloud"`
		Crio                           bool   `json:"crio"`
		Docker                         string `json:"docker"`
		DockerEnterprise               bool   `json:"dockerEnterprise"`
		HasPackageManager              bool   `json:"hasPackageManager"`
		K8SApiServer                   bool   `json:"k8sApiServer"`
		K8SControllerManager           bool   `json:"k8sControllerManager"`
		K8SEtcd                        bool   `json:"k8sEtcd"`
		K8SFederationApiServer         bool   `json:"k8sFederationApiServer"`
		K8SFederationControllerManager bool   `json:"k8sFederationControllerManager"`
		K8SKubelet                     bool   `json:"k8sKubelet"`
		K8SProxy                       bool   `json:"k8sProxy"`
		K8SScheduler                   bool   `json:"k8sScheduler"`
		Kubernetes                     string `json:"kubernetes"`
		Openshift                      bool   `json:"openshift"`
		OpenshiftVersion               string `json:"openshiftVersion"`
		OsDistro                       string `json:"osDistro"`
		Serverless                     bool   `json:"serverless"`
		SwarmManager                   bool   `json:"swarmManager"`
		SwarmNode                      bool   `json:"swarmNode"`
	} `json:"installedProducts"`
	Instances []struct {
		Host     string    `json:"host"`
		Image    string    `json:"image"`
		Modified time.Time `json:"modified"`
		Registry string    `json:"registry"`
		Repo     string    `json:"repo"`
		Tag      string    `json:"tag"`
	} `json:"instances"`
	IsARM64                   bool     `json:"isARM64"`
	K8SClusterAddr            string   `json:"k8sClusterAddr"`
	Labels                    []string `json:"labels"`
	Layers                    []string `json:"layers"`
	MissingDistroVulnCoverage bool     `json:"missingDistroVulnCoverage"`
	Namespaces                []string `json:"namespaces"`
	OsDistro                  string   `json:"osDistro"`
	OsDistroRelease           string   `json:"osDistroRelease"`
	OsDistroVersion           string   `json:"osDistroVersion"`
	PackageManager            bool     `json:"packageManager"`
	Packages                  []struct {
		Pkgs []struct {
			BinaryIdx  []int    `json:"binaryIdx"`
			BinaryPkgs []string `json:"binaryPkgs"`
			CveCount   int      `json:"cveCount"`
			Files      []struct {
				Md5    string `json:"md5"`
				Path   string `json:"path"`
				Sha1   string `json:"sha1"`
				Sha256 string `json:"sha256"`
			} `json:"files"`
			FunctionLayer string `json:"functionLayer"`
			LayerTime     int    `json:"layerTime"`
			License       string `json:"license"`
			Name          string `json:"name"`
			Path          string `json:"path"`
			Version       string `json:"version"`
		} `json:"pkgs"`
		PkgsType []string `json:"pkgsType"`
	} `json:"packages"`
	PullDuration      int      `json:"pullDuration"`
	RegistryNamespace string   `json:"registryNamespace"`
	RepoDigests       []string `json:"repoDigests"`
	RepoTag           struct {
		Digest   string `json:"digest"`
		Id       string `json:"id"`
		Registry string `json:"registry"`
		Repo     string `json:"repo"`
		Tag      string `json:"tag"`
	} `json:"repoTag"`
	RhelRepos   []string `json:"rhelRepos"`
	RiskFactors struct {
		Property1 string `json:"property1"`
		Property2 string `json:"property2"`
	} `json:"riskFactors"`
	ScanDuration    int       `json:"scanDuration"`
	ScanID          int       `json:"scanID"`
	ScanTime        time.Time `json:"scanTime"`
	ScanVersion     string    `json:"scanVersion"`
	StartupBinaries []struct {
		Altered       bool     `json:"altered"`
		CveCount      int      `json:"cveCount"`
		Deps          []string `json:"deps"`
		FunctionLayer string   `json:"functionLayer"`
		Md5           string   `json:"md5"`
		MissingPkg    bool     `json:"missingPkg"`
		Name          string   `json:"name"`
		Path          string   `json:"path"`
		PkgRootDir    string   `json:"pkgRootDir"`
		Services      []string `json:"services"`
		Version       string   `json:"version"`
	} `json:"startupBinaries"`
	Stopped bool `json:"stopped"`
	Tags    []struct {
		Digest   string `json:"digest"`
		Id       string `json:"id"`
		Registry string `json:"registry"`
		Repo     string `json:"repo"`
		Tag      string `json:"tag"`
	} `json:"tags"`
	TopLayer    string `json:"topLayer"`
	TrustResult struct {
		Groups []struct {
			Id           string    `json:"_id"`
			Disabled     bool      `json:"disabled"`
			Images       []string  `json:"images"`
			Layers       []string  `json:"layers"`
			Modified     time.Time `json:"modified"`
			Name         string    `json:"name"`
			Notes        string    `json:"notes"`
			Owner        string    `json:"owner"`
			PreviousName string    `json:"previousName"`
		} `json:"groups"`
		HostsStatuses []struct {
			Host   string   `json:"host"`
			Status []string `json:"status"`
		} `json:"hostsStatuses"`
	} `json:"trustResult"`
	TrustStatus     []string `json:"trustStatus"`
	TwistlockImage  bool     `json:"twistlockImage"`
	Type            []string `json:"type"`
	Vulnerabilities []struct {
		ApplicableRules []string  `json:"applicableRules"`
		BinaryPkgs      []string  `json:"binaryPkgs"`
		Block           bool      `json:"block"`
		Cause           string    `json:"cause"`
		Cri             bool      `json:"cri"`
		Custom          bool      `json:"custom"`
		Cve             string    `json:"cve"`
		Cvss            int       `json:"cvss"`
		Description     string    `json:"description"`
		Discovered      time.Time `json:"discovered"`
		Exploit         []string  `json:"exploit"`
		FixDate         int       `json:"fixDate"`
		FixLink         string    `json:"fixLink"`
		FunctionLayer   string    `json:"functionLayer"`
		GracePeriodDays int       `json:"gracePeriodDays"`
		Id              int       `json:"id"`
		LayerTime       int       `json:"layerTime"`
		Link            string    `json:"link"`
		PackageName     string    `json:"packageName"`
		PackageVersion  string    `json:"packageVersion"`
		Published       int       `json:"published"`
		RiskFactors     struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"riskFactors"`
		Severity     string     `json:"severity"`
		Status       string     `json:"status"`
		Templates    [][]string `json:"templates"`
		Text         string     `json:"text"`
		Title        string     `json:"title"`
		Twistlock    bool       `json:"twistlock"`
		Type         []string   `json:"type"`
		VecStr       string     `json:"vecStr"`
		VulnTagInfos []struct {
			Color   string `json:"color"`
			Comment string `json:"comment"`
			Name    string `json:"name"`
		} `json:"vulnTagInfos"`
	} `json:"vulnerabilities"`
	VulnerabilitiesCount      int `json:"vulnerabilitiesCount"`
	VulnerabilityDistribution struct {
		Critical int `json:"critical"`
		High     int `json:"high"`
		Low      int `json:"low"`
		Medium   int `json:"medium"`
		Total    int `json:"total"`
	} `json:"vulnerabilityDistribution"`
	VulnerabilityRiskScore int `json:"vulnerabilityRiskScore"`
	WildFireUsage          struct {
		Bytes   int `json:"bytes"`
		Queries int `json:"queries"`
		Uploads int `json:"uploads"`
	} `json:"wildFireUsage"`
}
