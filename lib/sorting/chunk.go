package sorting

import (
	"fmt"
	"image"
	"math/rand"
	"sort"
)

type chunk struct {
	begin int
	size  int
}

func chunkSize(min, max int) int {
	if min > max {
		// should never happen, because input will be validated before (with graceful shutdown)
		panic(fmt.Errorf("minimum chunk size can't be bigger than maximum chunk size"))
	}
	return rand.Intn(max-min) + min
}

type rgbaPixel [4]uint8
type rgbaPixels []rgbaPixel

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

	sort.Sort(byGrayscale(pixels))

	for i, pixel := range pixels {
		data[i*4+0] = pixel[0]
		data[i*4+1] = pixel[1]
		data[i*4+2] = pixel[2]
		data[i*4+3] = pixel[3]
	}
}
