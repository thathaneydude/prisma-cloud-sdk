# CWP Examples
Examples coming soon...
* Connecting to Self-Hosted Instance
* Listing Images

### Self-Hosted Instance 
```go
client, err := cwpp.NewCwppClient(
    "console.palo.com",
    "v22.01",
    false,
    "https")
authResp, err = client.Authenticate("API-KEY-ID", "API-SECRET")
```
### Listing Images
```go
query := cwpp.ImageQuery{
    Offset:      "0",
    Limit:       "10",
    Collections: []string{"collection_one"},
}
images, err := prisma_client.Cwpp.ListImages(query)
fmt.Println(fmt.Sprintf("First 10 images in collection_one: %v", images))
```
