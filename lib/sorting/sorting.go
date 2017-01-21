package sorting

import (
	"image"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func SortRGBA(rgba *image.RGBA, min, max int) {
	chunks := calculateChunks(rgba.Bounds(), min, max)

	for _, chunk := range chunks {
		chunk.sort(rgba)
	}
}

func calculateChunks(bounds image.Rectangle, min, max int) []chunk {
	chunks := []chunk{}

	size := bounds.Max.Sub(bounds.Min)
	width := size.X
	height := size.Y

	for y := 0; y < height; y++ {
		x := 0
		for x < width {
			size := chunkSize(min, max)
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
