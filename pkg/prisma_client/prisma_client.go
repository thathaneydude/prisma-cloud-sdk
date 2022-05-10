package prisma_client

import (
	"prisma-cloud-sdk/pkg/cspm"
	"prisma-cloud-sdk/pkg/cwpp"
)

type PrismaCloudClient struct {
	cwppBaseUrl    string
	cspmBaseUrl    string
	cwppApiVersion string
	sslVerify      bool
	schema         string
	Cwpp           *cwpp.CwppClient
	Cspm           *cspm.CspmClient
}
