package prisma

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thathaneydude/prisma-cloud-sdk/cs"
	cspm2 "github.com/thathaneydude/prisma-cloud-sdk/cspm"
	"github.com/thathaneydude/prisma-cloud-sdk/cwpp"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/client"
)

// NewPrismaCloudClient creates a Prisma Cloud client with the given Options for interacting with CSPM, CWPP, and CS APIs.
// Generally API Url, Username (access key) and Password (secret key) should be all that needs to be defined however
// other options are available. If you are a multi-tenant user and the body parameters of your login request include
// a user name and password instead of an access key ID and secret key.
// You will also need to provide the prismaId or the customerName. Specifying just the prismaId is preferred, but specifying
// customerName is an acceptable alternative. Your prismaId is available from the license information in the
// Prisma Cloud console. It's unnecessary to specify both prismaId and customerName, but if you do specify both,
// the parameters must indicate the same tenant.
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

	cspmClient, err := cspm2.NewCSPMClient(&cspm2.ClientOptions{
		ApiUrl:     o.ApiUrl,
		SslVerify:  o.SslVerify,
		Schema:     o.Schema,
		MaxRetries: o.MaxRetries,
	})
	if err != nil {
		logrus.Errorf(err.Error())
		return nil, err
	}
	cspmClient.OverwriteBaseClient(baseClient)
	cspmLoginReq := &cspm2.LoginRequest{
		Username: o.Username,
		Password: o.Password,
	}
	// Prisma ID / Customer Name should be
	if o.PrismaId != "" {
		cspmLoginReq.PrismaId = o.PrismaId
	}
	if o.CustomerName != "" {
		cspmLoginReq.CustomerName = o.CustomerName
	}

	_, err = cspmClient.Login(cspmLoginReq)
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

	cwppClient.OverwriteBaseClient(baseClient)

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
	CustomerName   string
	PrismaId       string
}
