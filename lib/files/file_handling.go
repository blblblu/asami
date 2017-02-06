package files

import (
	"image"
	_ "image/jpeg" // to be able to import jpeg images
	"image/png"
	"os"
)

func LoadImage(filename string) (image.Image, error) {
	inFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	img, _, err := image.Decode(inFile)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func SaveImage(filename string, img image.Image) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	png.Encode(outFile, img)

	return nil
}
