package prisma

import (
	"github.com/thathaneydude/prisma-cloud-sdk/cs"
	"github.com/thathaneydude/prisma-cloud-sdk/cspm"
	"github.com/thathaneydude/prisma-cloud-sdk/cwpp"
)

type PrismaCloudClient struct {
	cwppBaseUrl    string
	cspmBaseUrl    string
	cwppApiVersion string
	sslVerify      bool
	schema         string
	Cwpp           *cwpp.CwppClient
	Cspm           *cspm.CspmClient
	Cs             *cs.CsClient
}
