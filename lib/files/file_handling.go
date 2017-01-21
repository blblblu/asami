package files

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

func LoadImage(filename string) (*image.RGBA, error) {
	inFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	img, _, err := image.Decode(inFile)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	return rgba, nil
}

func SaveImage(filename string, rgba *image.RGBA) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	png.Encode(outFile, rgba)

	return nil
}
