package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/blblblu/asami/lib/commands"
	"github.com/urfave/cli"
)

var (
	version = "master"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	app := cli.App{
		Name:    "asami",
		Usage:   "simple image corruptor",
		Version: version,
		Commands: []*cli.Command{
			commands.NewSortCommand(),
		},
	}

	app.Run(os.Args)
}
