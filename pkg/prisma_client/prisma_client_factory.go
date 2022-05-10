package prisma_client

import (
	"PrismaCloud/pkg"
	"PrismaCloud/pkg/client"
	"PrismaCloud/pkg/constants"
	"PrismaCloud/pkg/cspm"
	"PrismaCloud/pkg/cwpp"
	"fmt"
	"github.com/sirupsen/logrus"
)

func NewDefaultPrismaCloudClient(apiUrl string, username string, password string, sslVerify bool) (*PrismaCloudClient, error) {
	return NewPrismaCloudClient(apiUrl, constants.DefaultSchema, username, password, constants.DefaultCwppApiVersion, constants.DefaultMaxRetries, sslVerify)
}

func NewPrismaCloudClient(apiUrl string, schema string, username string, password string, cwppApiVersion string, maxRetries int, sslVerify bool) (*PrismaCloudClient, error) {
	baseClient := client.NewBaseClient(sslVerify, maxRetries, schema)

	cspmClient, err := cspm.NewCSPMClient(apiUrl, sslVerify, schema, maxRetries)
	if err != nil {
		logrus.Errorf(err.Error())
		return nil, err
	}
	cspmClient.BaseClient = *baseClient

	_, err = cspmClient.Login(username, password)
	resp, err := cspmClient.GetMetaInfo()
	if err != nil {
		return nil, err
	}

	cwppClient, err := cwpp.NewCwppClient(resp.TwistlockUrl, cwppApiVersion, sslVerify, schema)
	if err != nil {
		return nil, &pkg.GenericError{Msg: fmt.Sprintf("Failed to initialize CWPP client using meta_info: %v", err)}
	}

	cwppClient.BaseClient = *baseClient

	return &PrismaCloudClient{
		cwppBaseUrl:    resp.TwistlockUrl,
		cwppApiVersion: cwppApiVersion,
		cspmBaseUrl:    apiUrl,
		Cwpp:           cwppClient,
		Cspm:           cspmClient,
	}, nil
}