package prisma

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thathaneydude/prisma-cloud-sdk/cwpp"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/client"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/cs"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/cspm"
)

// NewPrismaCloudClient creates a CWPP, CSPM, and Code Security client with the given credentials. This factory
// allows full configuration of the PrismaCloudClient.
func NewPrismaCloudClient(o *Options) (*PrismaCloudClient, error) {
	if o.schema == "" {
		o.schema = internal.DefaultSchema
	}

	if o.maxRetries == 0 {
		o.maxRetries = internal.DefaultMaxRetries
	}

	if o.cwppApiVersion == "" {
		o.cwppApiVersion = internal.DefaultCwppApiVersion
	}

	if o.apiUrl == "" || o.username == "" || o.password == "" {
		return nil, &internal.GenericError{Msg: "Required Client fields not correctly populated. API Url, " +
			"Username (Access Key ID), and Password (Secret Access Key) must be set"}
	}

	// Re-use the same base http client for each CSPM, CWPP, and CS clients
	baseClient := client.NewBaseClient(o.sslVerify, o.maxRetries, o.schema)

	cspmClient, err := cspm.NewCSPMClient(&cspm.ClientOptions{
		ApiUrl:     o.apiUrl,
		SslVerify:  o.sslVerify,
		Schema:     o.schema,
		MaxRetries: o.maxRetries,
	})
	if err != nil {
		logrus.Errorf(err.Error())
		return nil, err
	}
	cspmClient.BaseClient = *baseClient

	_, err = cspmClient.Login(o.username, o.password)
	if err != nil {
		return nil, err
	}

	// Fetch the meta_info for which contains the twistlock URL
	resp, err := cspmClient.GetMetaInfo()
	if err != nil {
		return nil, err
	}

	cwppClient, err := cwpp.NewCwppClient(&cwpp.ClientOptions{
		ConsoleUrl: resp.TwistlockUrl,
		ApiVersion: o.cwppApiVersion,
		SslVerify:  o.sslVerify,
		MaxRetries: o.maxRetries,
		Schema:     o.schema,
	})
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to initialize CWPP client using meta_info: %v", err)}
	}

	cwppClient.BaseClient = *baseClient

	// Code Security is basically just a CSPM client but with specifically exported functions relevant to Code Security
	csClient := cs.CsClientWithCspmInjected(cspmClient)

	return &PrismaCloudClient{
		cwppBaseUrl:    resp.TwistlockUrl,
		cwppApiVersion: o.cwppApiVersion,
		cspmBaseUrl:    o.apiUrl,
		Cwpp:           cwppClient,
		Cspm:           cspmClient,
		Cs:             csClient,
	}, nil
}

type Options struct {
	apiUrl         string
	schema         string
	username       string
	password       string
	cwppApiVersion string
	maxRetries     int
	sslVerify      bool
}
