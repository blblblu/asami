package main

import (
	"os"

	"github.com/urfave/cli"
)

var version = "master"

/*type Pixels [][]uint8

func (p Pixels) Len() int { return len(p) }
func (p Pixels) Swap(i, j int) {
	//p[i], p[j] = p[j], p[i]
	for k := 0; k < 2; k++ {
		p[i][k], p[j][k] = p[j][k], p[i][k]
	}
}
func (p Pixels) Less(i, j int) bool {
	sumI := 0
	sumJ := 0
	for x := 0; x < 4; x++ {
		sumI += int(p[i][x])
		sumJ += int(p[j][x])
	}
	return sumI < sumJ
}
func handleRequest(input string, output string) error {
	inputImage, _ := imaging.Open(input)
	x := inputImage.Bounds().Dx()
	y := inputImage.Bounds().Dy()
	outputImage := imaging.New(x, y, color.NRGBA{0, 0, 0, 0})
	//outputImage = imaging.PasteCenter(outputImage, inputImage)
	src := inputImage.(*image.NRGBA)
	dst := outputImage
	const chunkSize = 200
	count := x * y * 4
	for i := 0; i < count/(chunkSize+1); i++ {
		pixels := make([][]uint8, chunkSize)
		for j := range pixels {
			pixels[j] = src.Pix[i*chunkSize+j*4 : i*chunkSize+j*4+4]
		}
		sort.Sort(Pixels(pixels))
		for j := range pixels {
			dst.Pix[i*chunkSize+j*4] = pixels[j][0]
			dst.Pix[i*chunkSize+j*4+1] = pixels[j][1]
			dst.Pix[i*chunkSize+j*4+2] = pixels[j][2]
			dst.Pix[i*chunkSize+j*4+3] = pixels[j][3]
		}
		//dst.Pix[i*chunkSize+j] = src.Pix[i*chunkSize+j]
	}
	if err := imaging.Save(outputImage, output); err != nil {
		return err
	}
	return nil
}*/

func main() {
	app := &cli.App{
		Name:    "asami",
		Usage:   "simple pixel sorter",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "input",
				Usage:   "the input file `path` to use",
				Aliases: []string{"i"},
			},
			&cli.StringFlag{
				Name:    "output",
				Usage:   "the output file `path` to use",
				Aliases: []string{"o"},
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	app.Run(os.Args)
}
