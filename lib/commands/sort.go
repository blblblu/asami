package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/blblblu/asami/lib/files"
	"github.com/blblblu/asami/lib/sorting"
	"github.com/urfave/cli"
)

type sortingArgs struct {
	input  string
	output string
	min    int
	max    int
}

func NewSortCommand() *cli.Command {
	args := sortingArgs{}

	return &cli.Command{
		Name:  "sort",
		Usage: "simple brute sorting",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "input",
				Usage:       "the input file `path` to use",
				Aliases:     []string{"i"},
				Destination: &args.input,
			},
			&cli.StringFlag{
				Name:        "output",
				Usage:       "the output file `path` to use, must be a png file",
				Aliases:     []string{"o"},
				Destination: &args.output,
			},
			&cli.IntFlag{
				Name:        "min",
				Usage:       "the minimum chunk `size`",
				Value:       32,
				Destination: &args.min,
			},
			&cli.IntFlag{
				Name:        "max",
				Usage:       "the maximum chunk `size`",
				Value:       64,
				Destination: &args.max,
			},
		},
		Action: func(ctx *cli.Context) error {
			if err := checkSortingArgs(ctx, &args); err != nil {
				fmt.Fprintf(os.Stderr, "%s\n\n", err)
				cli.ShowCommandHelp(ctx, "sort")
				return cli.Exit("", 1)
			}

			rgba, err := files.LoadImage(args.input)
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed reading input file: %s", err), 2)
			}

			sorting.SortRGBA(rgba, args.min, args.max)

			err = files.SaveImage(args.output, rgba)
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed writing output image to file: %s", err), 3)
			}

			return nil
		},
	}
}

func checkSortingArgs(ctx *cli.Context, args *sortingArgs) error {
	errors := []string{}
	if args.input == "" {
		errors = append(errors, "input file path must be set")
	}
	if args.output == "" {
		errors = append(errors, "output file path must be set")
	}
	if args.min < 1 {
		errors = append(errors, "minimum chunk size must be at least 1")
	}
	if args.max < args.min {
		errors = append(errors, "maximum chunk size must be at least as big as the minimum chunk size")
	}

	if len(errors) > 0 {
		errorMessage := strings.Join(errors, "\n")
		return fmt.Errorf(errorMessage)
	}

	return nil
}
