package prisma_client

import (
	"prisma-cloud-sdk/cspm"
	"prisma-cloud-sdk/cwpp"
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
