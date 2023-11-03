package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	tnp "github.com/ProfChaos/torrent-name-parser"
)

var (
	flags = flag.NewFlagSet("tnp", flag.ExitOnError)
	debug = flags.Bool("debug", false, "Debug the parser")
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tnp <torrent name> [-debug]")
		return
	}

	err := flags.Parse(os.Args[2:])
	if err != nil {
		panic(err)
	}

	if debug != nil && *debug {
		tnp.DebugParser(os.Args[1])
		return
	}

	torrent, err := tnp.ParseName(os.Args[1])
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(torrent)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
