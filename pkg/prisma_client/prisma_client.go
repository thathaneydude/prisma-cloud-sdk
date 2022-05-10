package prisma_client

import (
	"PrismaCloud/pkg/cspm"
	"PrismaCloud/pkg/cwpp"
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
