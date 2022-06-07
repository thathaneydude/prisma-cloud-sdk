package prisma

import (
	"github.com/thathaneydude/prisma-cloud-sdk/cwpp"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/cs"
	"github.com/thathaneydude/prisma-cloud-sdk/internal/cspm"
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
