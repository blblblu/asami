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

type rgbaPixels []uint8

func (chunk *chunk) sort(rgba *image.RGBA) {
	beginIndex := chunk.begin * 4
	endIndex := beginIndex + chunk.size*4
	data := rgba.Pix[beginIndex:endIndex]

	sort.Sort(byGrayscale(data))
}
