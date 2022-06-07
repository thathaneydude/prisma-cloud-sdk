# Prisma Cloud SDK 
prisma-cloud-sdk is the unofficial Prisma Cloud SDK for the Go programming language.

# Installation
Use `go get` to retrieve the SDK and addit to your `GOPATH` workspace, or project's Go module dependencies.
```
$ go get -u https://github.com/thathaneydude/prisma-cloud-sdk
```

## Dependencies


# Creating a Client
## Creating a New Default Prisma Cloud Client
```go
// Your CWPP console will be automatically determined
prisma_client, err := prisma.NewDefaultPrismaCloudClient(
	"api.prismacloud.io", // API URL to connect to. More info here: https://prisma.pan.dev/api/cloud/api-urls 
	"API-KEY-ID", // https://docs.paloaltonetworks.com/prisma/prisma-cloud/prisma-cloud-admin/manage-prisma-cloud-administrators/create-access-keys
	"API-SECRET", 
	false) // Setting to true will use the local system default cert pool unless the env vars SSL_CERT_FILE and SSL_CERT_DIR are set		
```

## Create a Custom Client
```go
// Your CWPP console will be automatically determined
prisma_client, err := prisma.NewPrismaCloudClient(
    "api.prismacloud.io", // API URL to connect to. More info here: https://prisma.pan.dev/api/cloud/api-urls
    "https", // typically, only changed for testing
    "API-KEY-ID", // https://docs.paloaltonetworks.com/prisma/prisma-cloud/prisma-cloud-admin/manage-prisma-cloud-administrators/create-access-keys
    "API-SECRET", 
    "v22.01", // Only current version and current version -1 will be available. Setting to "1" will always use latest
    5, // Number of times the SDK will retry a request if a 429 is returned
    false)
```

## Examples
