package sorting

import (
	"image"
	"image/draw"
	"math/rand"
	"time"

	"github.com/disintegration/imaging"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func SortImage(img image.Image, min, max int, inverted bool) image.Image {
	rgba := image.NewRGBA(img.Bounds())

	if !inverted {
		draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	} else {
		invertedImg := imaging.Invert(img)
		draw.Draw(rgba, rgba.Bounds(), invertedImg, image.Point{0, 0}, draw.Src)
	}

	chunks := calculateChunks(rgba.Bounds(), min, max)

	for _, chunk := range chunks {
		chunk.sort(rgba)
	}

	if !inverted {
		return rgba
	}
	return imaging.Invert(rgba)
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
