package main

import (
	"flag"
	"fmt"
	"github.com/jeffwatts/war"
)

func main() {
	var logVerbose bool
	flag.BoolVar(&logVerbose, "v", false, "enable verbose logging")
	flag.Parse()

	if logVerbose {
		fmt.Println("creating War game, verbose logging enabled")
	}

	game := war.NewGame(logVerbose)
	game.Play()
}
