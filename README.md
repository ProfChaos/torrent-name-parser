# Torrent Name Parser

Torrent name parser in Go

Do feel free to create issues with sane names that fail to parse properly

Have done a little clean up and increased the test coverage to ~95%

## Usage

```sh
go get github.com/ProfChaos/torrent-name-parser
```

```go
package main

import (
  "fmt"

  tnp "github.com/ProfChaos/torrent-name-parser"
)

func main() {
  torrent, err := tnp.ParseName("blade.runner.2049.2017.2160p.uhd.bluray.x265-terminal.mkv")
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("%+#v\n", torrent)
  fmt.Println(torrent.Title)
}
```

Output:

```sh
torrentparser.Torrent{Title:"blade runner 2049", AlternativeTitle:"", ContentType:1, Year:2017, Resolution:"4k", Extended:false, Unrated:false, Proper:false, Repack:false, Convert:false, Hardcoded:false, Retail:false, Remastered:false, Region:"", Container:"mkv", Source:"bluray", Codec:"x265", Audio:"", Group:"terminal", Season:-1, Episode:0, Language:"", Hdr:false, ColorDepth:"", Date:""}
blade runner 2049
```
