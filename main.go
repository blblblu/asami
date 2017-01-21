package main

import (
	"image"
	"image/draw"
	"image/png"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var (
	version = "master"
	app     *cli.App
	args    arguments
)

type arguments struct {
	input  string
	output string
	min    int
	max    int
}

func init() {
	app = &cli.App{
		Name:    "asami",
		Usage:   "pixel sorter using simple brute sorting",
		Version: version,
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
		Action: func(c *cli.Context) error {
			if err := checkInput(c); err != nil {
				return err
			}

			return nil
		},
	}
}

func checkInput(c *cli.Context) error {
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
		return cli.Exit(errorMessage, 1)
	}

	image, err := loadImage(args.input)
	if err != nil {
		return err
	}

	sort(image)

	err = saveImage(args.output, image)
	if err != nil {
		return err
	}

	return nil
}

func loadImage(filename string) (*image.RGBA, error) {
	inFile, err := os.Open(filename)
	if err != nil {
		return nil, cli.Exit("failed reading input file", 2)
	}
	defer inFile.Close()

	img, _, err := image.Decode(inFile)
	if err != nil {
		return nil, cli.Exit("failed decoding input file", 3)
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	return rgba, nil
}

func saveImage(filename string, image *image.RGBA) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return cli.Exit("failed writing output image to file", 4)
	}
	defer outFile.Close()

	png.Encode(outFile, image)

	return nil
}

func sort(img *image.RGBA) {

}

func main() {
	app.Run(os.Args)
}
