# Zaya Go Client

A Go client library for the Zaya.io link shortener API.

## Installation

```bash
go get github.com/abtinokhovat/zaya-go
```

## Usage

First, you need to get your API key from Zaya.io.

### Creating a Client

```go
import "github.com/abtinokhovat/zaya-go"

client := zaya.NewClient("your-api-key")

// Optional: Configure timeout
client.WithTimeout(10 * time.Second)

// Optional: Configure custom base URL
client.WithBaseURL("https://custom-zaya-instance.com/api/v1")
```

### Links

```go
// List links
links, err := client.ListLinks(zaya.LinkListParams{
    Search: "example",
    Sort:   "desc",
})

// Create a link
link, err := client.CreateLink("https://example.com", zaya.CreateLinkParams{
    Alias:      "example",
    Password:   "secret",
    Public:     true,
})

// Get link details
link, err := client.GetLink(1)

// Update a link
link, err := client.UpdateLink(1, zaya.CreateLinkParams{
    Description: "Updated description",
})

// Delete a link
err := client.DeleteLink(1)
```

### Spaces

```go
// List spaces
spaces, err := client.ListSpaces(zaya.SpaceListParams{
    Search: "example",
    Sort:   "desc",
})

// Create a space
space, err := client.CreateSpace(zaya.CreateSpaceParams{
    Name:  "Example Space",
    Color: "#FF0000",
})

// Get space details
space, err := client.GetSpace(1)

// Update a space
space, err := client.UpdateSpace(1, zaya.CreateSpaceParams{
    Name:  "Updated Space",
    Color: "#00FF00",
})

// Delete a space
err := client.DeleteSpace(1)
```

### Domains

```go
// List domains
domains, err := client.ListDomains(zaya.DomainListParams{
    Search: "example",
    Sort:   "desc",
})

// Create a domain
domain, err := client.CreateDomain(zaya.CreateDomainParams{
    Name:         "example.com",
    IndexPage:    "https://example.com/index",
    NotFoundPage: "https://example.com/404",
})

// Get domain details
domain, err := client.GetDomain(1)

// Update a domain
domain, err := client.UpdateDomain(1, zaya.CreateDomainParams{
    IndexPage:    "https://example.com/new-index",
    NotFoundPage: "https://example.com/new-404",
})

// Delete a domain
err := client.DeleteDomain(1)
```

### Statistics

```go
// Get total statistics
stats, err := client.GetTotalStats(1, zaya.StatsParams{
    From: time.Now().AddDate(0, -1, 0),
    To:   time.Now(),
})

// Get click statistics
clicks, err := client.GetClickStats(1, zaya.StatsParams{
    From: time.Now().AddDate(0, -1, 0),
    To:   time.Now(),
})

// Get referrer statistics
referrers, err := client.GetReferrerStats(1, zaya.StatsParams{
    From: time.Now().AddDate(0, -1, 0),
    To:   time.Now(),
})

// Get country statistics
countries, err := client.GetCountryStats(1, zaya.StatsParams{
    From: time.Now().AddDate(0, -1, 0),
    To:   time.Now(),
})

// Get language statistics
languages, err := client.GetLanguageStats(1, zaya.StatsParams{
    From: time.Now().AddDate(0, -1, 0),
    To:   time.Now(),
})

// Get browser statistics
browsers, err := client.GetBrowserStats(1, zaya.StatsParams{
    From: time.Now().AddDate(0, -1, 0),
    To:   time.Now(),
})

// Get device statistics
devices, err := client.GetDeviceStats(1, zaya.StatsParams{
    From: time.Now().AddDate(0, -1, 0),
    To:   time.Now(),
})

// Get operating system statistics
os, err := client.GetOperatingSystemStats(1, zaya.StatsParams{
    From: time.Now().AddDate(0, -1, 0),
    To:   time.Now(),
})
```

### Account

```go
// Get account details
account, err := client.GetAccount()
```

## Error Handling

All methods return an error that should be checked:

```go
link, err := client.CreateLink("https://example.com", zaya.CreateLinkParams{})
if err != nil {
    // Handle error
    log.Fatal(err)
}
```

## License

MIT
