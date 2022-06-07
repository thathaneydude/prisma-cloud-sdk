# Prisma Cloud SDK 
prisma-cloud-sdk is the unofficial Prisma Cloud SDK for the Go programming language.

# Installation
Use `go get` to retrieve the SDK and addit to your `GOPATH` workspace, or project's Go module dependencies.
```
$ go get -u https://github.com/thathaneydude/prisma-cloud-sdk
```

## Dependencies


## Create a Prisma Cloud Client
```go
// Your CWPP console will be automatically determined
prisma_client, err := prisma.NewPrismaCloudClient((&prisma.Options{
    ApiUrl:    "api.prismacloud.io", // API URL to connect to. More info here: https://prisma.pan.dev/api/cloud/api-urls
    Username:  "API-KEY-ID",         // https://docs.paloaltonetworks.com/prisma/prisma-cloud/prisma-cloud-admin/manage-prisma-cloud-administrators/create-access-keys
    Password:  "API-SECRET",
    SslVerify: false,
})
```