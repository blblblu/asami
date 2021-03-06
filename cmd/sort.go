package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/blblblu/asami/lib/files"
	"github.com/blblblu/asami/lib/sorting"
	"github.com/spf13/cobra"
)

var sortingArgs struct {
	input    string
	output   string
	min      int
	max      int
	inverted bool
}

var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "simple brute pixel sorting",
	Long:  `simple brute pixel sorting`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkSortingArgs(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n\n", err)
			cmd.Help()
			os.Exit(1)
			return
		}

		img, err := files.LoadImage(sortingArgs.input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed reading input file: %s", err)
			os.Exit(2)
			return
		}

		sorted := sorting.SortImage(img, sortingArgs.min, sortingArgs.max, sortingArgs.inverted)

		err = files.SaveImage(sortingArgs.output, sorted)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed writing output image to file: %s", err)
			os.Exit(3)
			return
		}
	},
}

func checkSortingArgs() error {
	errors := []string{}
	if sortingArgs.input == "" {
		errors = append(errors, "input file path must be set")
	}
	if sortingArgs.output == "" {
		errors = append(errors, "output file path must be set")
	}
	if sortingArgs.min < 1 {
		errors = append(errors, "minimum chunk size must be at least 1")
	}
	if sortingArgs.max < sortingArgs.min {
		errors = append(errors, "maximum chunk size must be at least as big as the minimum chunk size")
	}

	if len(errors) > 0 {
		errorMessage := strings.Join(errors, "\n")
		return fmt.Errorf(errorMessage)
	}

	return nil
}

func init() {
	RootCmd.AddCommand(sortCmd)

	sortCmd.Flags().StringVarP(&sortingArgs.input, "input", "i", "", "the input file path to use")
	sortCmd.Flags().StringVarP(&sortingArgs.output, "output", "o", "", "the output file path to use, must be a png file")
	sortCmd.Flags().IntVar(&sortingArgs.min, "min", 32, "the minimum chunk size to use")
	sortCmd.Flags().IntVar(&sortingArgs.max, "max", 64, "the maximum chunk size to use")
	sortCmd.Flags().BoolVar(&sortingArgs.inverted, "inverted", false, "inverts the sorting direction")
}
