package main

import (
	"os"

	"github.com/blblblu/asami/lib/commands"
	"github.com/urfave/cli"
)

var (
	version     = "master"
	app         *cli.App
	sortingArgs commands.SortingArgs
)

func init() {
	app = &cli.App{
		Name:    "asami",
		Usage:   "simple image corruptor",
		Version: version,
		Commands: []*cli.Command{
			commands.NewSortCommand(&sortingArgs),
		},
	}
}

func main() {
	app.Run(os.Args)
}
