package sorting

import (
	"image"
	"image/draw"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func SortImage(img image.Image, min, max int) image.Image {
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	chunks := calculateChunks(rgba.Bounds(), min, max)

	for _, chunk := range chunks {
		chunk.sort(rgba)
	}

	return rgba
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
