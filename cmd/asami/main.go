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
		Name:  "asami",
		Usage: "simple image corruptor",
		Authors: []*cli.Author{
			{Name: "Sebastian Schulz", Email: "mail@sesc.me"},
		},
		Version: version,
		Commands: []*cli.Command{
			commands.NewSortCommand(),
		},
	}

	app.Run(os.Args)
}
