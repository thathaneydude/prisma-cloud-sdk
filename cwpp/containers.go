package cwpp

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/schema"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
)

const containersEndpoint = "/containers"

func (c *CwppClient) ListContainers(query ContainerQuery) ([]Container, error) {
	var encoder = schema.NewEncoder()
	params := url.Values{}
	err := encoder.Encode(query, params)
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to decode Host query provided: %v", err)}
	}
	var containersResponse []Container
	err = c.getWithResponseInterface(containersEndpoint, params, &containersResponse)
	if err != nil {
		return nil, err
	}
	return containersResponse, nil
}

type ContainerQuery struct {
	Offset            string   `schema:"offset,omitempty"`
	Limit             string   `schema:"limit,omitempty"`
	Search            string   `schema:"search,omitempty"`
	Sort              string   `schema:"sort,omitempty"`
	Reverse           bool     `schema:"reverse,omitempty"`
	Collections       []string `schema:"collections,omitempty"`
	AccountIDs        []string `schema:"accountIDs,omitempty"`
	Fields            []string `schema:"fields,omitempty"`
	Hostname          []string `schema:"hostname,omitempty"`
	Image             []string `schema:"image,omitempty"`
	ImageId           []string `schema:"imageId,omitempty"`
	Id                []string `schema:"id,omitempty"`
	profileID         []string `schema:"profileId,omitempty"`
	Namespaces        []string `schema:"namespaces,omitempty"`
	FirewallSupported bool     `schema:"firewallSupported,omitempty"`
	Clusters          []string `schema:"clusters,omitempty"`
	ComplianceIds     []int    `schema:"complianceIds,omitempty"`
}

type Container struct {
	ID                 string   `json:"_id"`
	Collections        []string `json:"collections"`
	FirewallProtection struct {
		Enabled              bool   `json:"enabled"`
		OutOfBandMode        string `json:"outOfBandMode"`
		Ports                []int  `json:"ports"`
		Supported            bool   `json:"supported"`
		TLSPorts             []int  `json:"tlsPorts"`
		UnprotectedProcesses []struct {
			Port    int    `json:"port"`
			Process string `json:"process"`
			TLS     bool   `json:"tls"`
		} `json:"unprotectedProcesses"`
	} `json:"firewallProtection"`
	Hostname string `json:"hostname"`
	Info     struct {
		AllCompliance struct {
			Compliance []struct {
				ApplicableRules []string          `json:"applicableRules"`
				BinaryPkgs      []string          `json:"binaryPkgs"`
				Block           bool              `json:"block"`
				Cause           string            `json:"cause"`
				Cri             bool              `json:"cri"`
				Custom          bool              `json:"custom"`
				Cve             string            `json:"cve"`
				Cvss            int               `json:"cvss"`
				Description     string            `json:"description"`
				Discovered      time.Time         `json:"discovered"`
				Exploit         string            `json:"exploit"`
				FixDate         int               `json:"fixDate"`
				FixLink         string            `json:"fixLink"`
				FunctionLayer   string            `json:"functionLayer"`
				GracePeriodDays int               `json:"gracePeriodDays"`
				ID              int               `json:"id"`
				LayerTime       int               `json:"layerTime"`
				Link            string            `json:"link"`
				PackageName     string            `json:"packageName"`
				PackageVersion  string            `json:"packageVersion"`
				Published       int               `json:"published"`
				RiskFactors     map[string]string `json:"riskFactors"`
				Severity        string            `json:"severity"`
				Status          string            `json:"status"`
				Templates       []string          `json:"templates"`
				Text            string            `json:"text"`
				Title           string            `json:"title"`
				Twistlock       bool              `json:"twistlock"`
				Type            []string          `json:"type"`
				VecStr          string            `json:"vecStr"`
				VulnTagInfos    []struct {
					Color   string `json:"color"`
					Comment string `json:"comment"`
					Name    string `json:"name"`
				} `json:"vulnTagInfos"`
			} `json:"compliance"`
			Enabled bool `json:"enabled"`
		} `json:"allCompliance"`
		App           string `json:"app"`
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
			Name        string `json:"name"`
			Provider    string `json:"provider"`
			Region      string `json:"region"`
			ResourceID  string `json:"resourceID"`
			ResourceURL string `json:"resourceURL"`
			Type        string `json:"type"`
			VMID        string `json:"vmID"`
		} `json:"cloudMetadata"`
		Cluster                string `json:"cluster"`
		ComplianceDistribution struct {
			Critical int `json:"critical"`
			High     int `json:"high"`
			Low      int `json:"low"`
			Medium   int `json:"medium"`
			Total    int `json:"total"`
		} `json:"complianceDistribution"`
		ComplianceIssues []struct {
			ApplicableRules []string          `json:"applicableRules"`
			BinaryPkgs      []string          `json:"binaryPkgs"`
			Block           bool              `json:"block"`
			Cause           string            `json:"cause"`
			Cri             bool              `json:"cri"`
			Custom          bool              `json:"custom"`
			Cve             string            `json:"cve"`
			Cvss            int               `json:"cvss"`
			Description     string            `json:"description"`
			Discovered      time.Time         `json:"discovered"`
			Exploit         string            `json:"exploit"`
			FixDate         int               `json:"fixDate"`
			FixLink         string            `json:"fixLink"`
			FunctionLayer   string            `json:"functionLayer"`
			GracePeriodDays int               `json:"gracePeriodDays"`
			ID              int               `json:"id"`
			LayerTime       int               `json:"layerTime"`
			Link            string            `json:"link"`
			PackageName     string            `json:"packageName"`
			PackageVersion  string            `json:"packageVersion"`
			Published       int               `json:"published"`
			RiskFactors     map[string]string `json:"riskFactors"`
			Severity        string            `json:"severity"`
			Status          string            `json:"status"`
			Templates       []string          `json:"templates"`
			Text            string            `json:"text"`
			Title           string            `json:"title"`
			Twistlock       bool              `json:"twistlock"`
			Type            string            `json:"type"`
			VecStr          string            `json:"vecStr"`
			VulnTagInfos    []struct {
				Color   string `json:"color"`
				Comment string `json:"comment"`
				Name    string `json:"name"`
			} `json:"vulnTagInfos"`
		} `json:"complianceIssues"`
		ComplianceIssuesCount int `json:"complianceIssuesCount"`
		ComplianceRiskScore   int `json:"complianceRiskScore"`
		ExternalLabels        []struct {
			Key        string    `json:"key"`
			SourceName string    `json:"sourceName"`
			SourceType []string  `json:"sourceType"`
			Timestamp  time.Time `json:"timestamp"`
			Value      string    `json:"value"`
		} `json:"externalLabels"`
		ID                string `json:"id"`
		Image             string `json:"image"`
		ImageID           string `json:"imageID"`
		ImageName         string `json:"imageName"`
		Infra             bool   `json:"infra"`
		InstalledProducts struct {
			Agentless                      bool   `json:"agentless"`
			Apache                         string `json:"apache"`
			AwsCloud                       bool   `json:"awsCloud"`
			Crio                           bool   `json:"crio"`
			Docker                         string `json:"docker"`
			DockerEnterprise               bool   `json:"dockerEnterprise"`
			HasPackageManager              bool   `json:"hasPackageManager"`
			K8SAPIServer                   bool   `json:"k8sApiServer"`
			K8SControllerManager           bool   `json:"k8sControllerManager"`
			K8SEtcd                        bool   `json:"k8sEtcd"`
			K8SFederationAPIServer         bool   `json:"k8sFederationApiServer"`
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
		Labels    []string `json:"labels"`
		Name      string   `json:"name"`
		Namespace string   `json:"namespace"`
		Network   struct {
			Ports []struct {
				Container int    `json:"container"`
				Host      int    `json:"host"`
				HostIP    string `json:"hostIP"`
				Listening bool   `json:"listening"`
				Nat       bool   `json:"nat"`
			} `json:"ports"`
		} `json:"network"`
		NetworkSettings struct {
			IPAddress  string `json:"ipAddress"`
			MacAddress string `json:"macAddress"`
			Networks   []struct {
				IPAddress  string `json:"ipAddress"`
				MacAddress string `json:"macAddress"`
				Name       string `json:"name"`
			} `json:"networks"`
			Ports []struct {
				ContainerPort string `json:"containerPort"`
				HostIP        string `json:"hostIP"`
				HostPort      int    `json:"hostPort"`
			} `json:"ports"`
		} `json:"networkSettings"`
		Processes []struct {
			Name string `json:"name"`
		} `json:"processes"`
		ProfileID string    `json:"profileID"`
		SizeBytes int       `json:"sizeBytes"`
		StartTime time.Time `json:"startTime"`
	} `json:"info"`
	ScanTime time.Time `json:"scanTime"`
}
