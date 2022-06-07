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
	if o.Schema == "" {
		o.Schema = internal.DefaultSchema
	}

	if o.MaxRetries == 0 {
		o.MaxRetries = internal.DefaultMaxRetries
	}

	if o.CwppApiVersion == "" {
		o.CwppApiVersion = internal.DefaultCwppApiVersion
	}

	if o.ApiUrl == "" || o.Username == "" || o.Password == "" {
		return nil, &internal.GenericError{Msg: "Required Client fields not correctly populated. API Url, " +
			"Username (Access Key ID), and Password (Secret Access Key) must be set"}
	}

	// Re-use the same base http client for each CSPM, CWPP, and CS clients
	baseClient := client.NewBaseClient(o.SslVerify, o.MaxRetries, o.Schema)

	cspmClient, err := cspm.NewCSPMClient(&cspm.ClientOptions{
		ApiUrl:     o.ApiUrl,
		SslVerify:  o.SslVerify,
		Schema:     o.Schema,
		MaxRetries: o.MaxRetries,
	})
	if err != nil {
		logrus.Errorf(err.Error())
		return nil, err
	}
	cspmClient.BaseClient = *baseClient

	_, err = cspmClient.Login(o.Username, o.Password)
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
		ApiVersion: o.CwppApiVersion,
		SslVerify:  o.SslVerify,
		MaxRetries: o.MaxRetries,
		Schema:     o.Schema,
	})
	if err != nil {
		return nil, &internal.GenericError{Msg: fmt.Sprintf("Failed to initialize CWPP client using meta_info: %v", err)}
	}

	cwppClient.BaseClient = *baseClient

	// Code Security is basically just a CSPM client but with specifically exported functions relevant to Code Security
	csClient := cs.CsClientWithCspmInjected(cspmClient)

	return &PrismaCloudClient{
		cwppBaseUrl:    resp.TwistlockUrl,
		cwppApiVersion: o.CwppApiVersion,
		cspmBaseUrl:    o.ApiUrl,
		Cwpp:           cwppClient,
		Cspm:           cspmClient,
		Cs:             csClient,
	}, nil
}

type Options struct {
	ApiUrl         string
	Schema         string
	Username       string
	Password       string
	CwppApiVersion string
	MaxRetries     int
	SslVerify      bool
}
