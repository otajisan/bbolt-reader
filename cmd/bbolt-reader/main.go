package main

import (
	"github.com/otajisan/bbolt-reader/bbolt"
	"os"
)

var (
	version = "0.0.1"
)

func main() {
	app := bbolt.NewApp(version)
	app.Run(os.Args)
}
