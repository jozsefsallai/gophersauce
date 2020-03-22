# gophersauce

gophersauce is an API wrapper for SauceNAO written in Go.

```
go get github.com/jozsefsallai/gophersauce
```

## Basic Example

```go
package main

import (
  "fmt"
  "log"
  "os"

  "github.com/jozsefsallai/gophersauce"
)

func main() {
  client, err := gophersauce.NewClient(nil)
  if err != nil {
    log.Fatal(err)
  }

  response, err := client.FromURL("https://i.imgur.com/v6EiHyj.png")
  if err != nil {
    log.Fatal(err)
  }

  if response.Count() == 0 {
    fmt.Println("There are no results :(")
    os.Exit(0)
  }

  first := response.First()
  if first.IsPixiv() {
    fmt.Println("This result is from Pixiv!")
    fmt.Println("Here's the ID of the image on Pixiv:", first.Data.PixivID)
  } else {
    fmt.Println("Probably not from Pixiv!")
    fmt.Println("Here are the external URLs:", first.Data.ExternalURLs)
  }
}
```

## Currently Supported Providers:
  * Pixiv
  * IMDb
  * Bcy
  * AniDB
  * Pawoo
  * Seiga
  * Sankaku
  * Danbooru

## Documentation and Reference

https://pkg.go.dev/github.com/jozsefsallai/gophersauce

## Todo
  * Add more providers
  * Unit tests

## License

MIT.
