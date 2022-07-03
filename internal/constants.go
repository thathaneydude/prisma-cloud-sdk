package internal

import "net/http"

const (
	DefaultMaxRetries     = 3
	DefaultCwppApiVersion = "1"
	UserAgent             = "PrismaCloudGoSDK"
	AuthHeader            = "x-redlock-auth"
	SupportedAPIURLLink   = "https://prisma.pan.dev/api/cloud/api-urls"
	DefaultSchema         = "https"
	RequestIdHeader       = "X-Redlock-Request-Id"
)

var APIVersions = []string{"22.06", "22.01", "21.08", "1"}
var SupportedHttpMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPatch,
	http.MethodPut,
	http.MethodDelete,
}

var SupportedAPIURLs = []string{
	"api.prismacloud.io",
	"api2.prismacloud.io",
	"api3.prismacloud.io",
	"api4.prismacloud.io",
	"api.anz.prismacloud.io",
	"api.eu.prismacloud.io",
	"api2.eu.prismacloud.io",
	"api.gov.prismacloud.io",
	"api.prismacloud.cn",
	"api.ca.prismacloud.io",
	"api.sg.prismacloud.io",
	"api.uk.prismacloud.io",
	"api.ind.prismacloud.io",
}

var CloudTypes = []string{
	"ALL",
	"AWS",
	"AZURE",
	"GCP",
	"ALIBABA_CLOUD",
	"OCI",
}
