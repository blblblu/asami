package main

import (
	"image"
	"image/draw"
	_ "image/jpeg" // to be able to import jpeg images
	"image/png"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

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

			rgba, err := loadImage(args.input)
			if err != nil {
				return err
			}

			sortRGBA(rgba)

			err = saveImage(args.output, rgba)
			if err != nil {
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

func saveImage(filename string, rgba *image.RGBA) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return cli.Exit("failed writing output image to file", 4)
	}
	defer outFile.Close()

	png.Encode(outFile, rgba)

	return nil
}

type chunk struct {
	begin int
	size  int
}

func sortRGBA(rgba *image.RGBA) {
	chunks := calculateChunks(rgba.Bounds())

	for _, chunk := range chunks {
		chunk.sort(rgba)
	}
}

func calculateChunks(bounds image.Rectangle) []chunk {
	chunks := []chunk{}

	size := bounds.Max.Sub(bounds.Min)
	width := size.X
	height := size.Y

	rand.Seed(time.Now().UnixNano())

	for y := 0; y < height; y++ {
		x := 0
		for x < width {
			size := chunkSize()
			if size > width-x {
				size = width - x
			}
			begin := y*width + x
			chunks = append(chunks, chunk{begin, size})
			x += size
		}
	}

	return chunks
}

func chunkSize() int {
	return rand.Intn(args.max-args.min) + args.min
}

type rgbaPixel [4]uint8
type rgbaPixels []rgbaPixel

type byRed rgbaPixels

func (p byRed) Len() int           { return len(p) }
func (p byRed) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byRed) Less(i, j int) bool { return p[i][0] < p[j][0] }

type byGreen rgbaPixels

func (p byGreen) Len() int           { return len(p) }
func (p byGreen) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byGreen) Less(i, j int) bool { return p[i][1] < p[j][1] }

type byBlue rgbaPixels

func (p byBlue) Len() int           { return len(p) }
func (p byBlue) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byBlue) Less(i, j int) bool { return p[i][2] < p[j][2] }

type bySum rgbaPixels

func (p bySum) Len() int      { return len(p) }
func (p bySum) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p bySum) Less(i, j int) bool {
	return p[i].sum() < p[j].sum()
}

func (p *rgbaPixel) sum() uint32 {
	return uint32(p[0]) + uint32(p[1]) + uint32(p[2])
}

type byGrayscale rgbaPixels

func (p byGrayscale) Len() int      { return len(p) }
func (p byGrayscale) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p byGrayscale) Less(i, j int) bool {
	return p[i].gray() < p[j].gray()
}

func (p *rgbaPixel) gray() uint8 {
	return p[0]>>2 + p[1]>>1 + p[1]>>3 + p[2]>>3
}

func (chunk *chunk) sort(rgba *image.RGBA) {
	beginIndex := chunk.begin * 4
	endIndex := beginIndex + chunk.size*4
	data := rgba.Pix[beginIndex:endIndex]

	pixels := rgbaPixels{}

	for i := 0; i < len(data); i += 4 {
		pixel := rgbaPixel{
			data[i+0],
			data[i+1],
			data[i+2],
			data[i+3],
		}
		pixels = append(pixels, pixel)
	}

	sort.Sort(byBlue(pixels))

	for i, pixel := range pixels {
		data[i*4+0] = pixel[0]
		data[i*4+1] = pixel[1]
		data[i*4+2] = pixel[2]
		data[i*4+3] = pixel[3]
	}
}

func main() {
	app.Run(os.Args)
}
