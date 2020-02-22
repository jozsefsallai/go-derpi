# go-derpi

Derpibooru API client for the Go programming language.

## Getting Started

**Add go-derpi to your dependencies:**

```
go get -u github.com/jozsefsallai/go-derpi
```

**Instantiate a client:**

```go
package derpi

import (
  "fmt"

  "github.com/jozsefsallai/go-derpi"
)

func main() {
  client := derpi.Init()

  // If you use a different endpoint or a proxy:
  // client.SetURL("https://myderpi.org/api/v1/json")

  // If you have an authorization key:
  // client.SetKey("your_key")

  image := client.GetImage(2207060)
  // will return the details of https://derpibooru.org/2207060
  // JSON response: https://derpibooru.org/api/v1/json/images/2207060
}
```

## Documentation

WIP.

## License

MIT.
