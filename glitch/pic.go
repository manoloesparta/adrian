package glitch

import (
	"image"
	"os"
)

// GlitchyImage is for transforming the img
type GlitchyImage struct {
	path   string
	pixels []uint32
}

// NewGlitchyImage constructor
func NewGlitchyImage(path string) GlitchyImage {
	img := loadImage(path)
	pixels := []uint32{}

	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels = append(pixels, r, g, b, a)
		}
	}

	return GlitchyImage{
		path,
		pixels,
	}
}

func loadImage(path string) image.Image {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic("Image no found")
	}

	img, _, err := image.Decode(file)
	if err != nil {
		panic("Format of image not identified")
	}
	return img
}
