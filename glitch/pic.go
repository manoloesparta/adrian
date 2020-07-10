package glitch

import "io/ioutil"

// GlitchyImage is for transforming the img
type GlitchyImage struct {
	path   string
	pixels []byte
}

// NewGlitchyImage constructor
func NewGlitchyImage(path string) GlitchyImage {
	pixels, err := ioutil.ReadFile(path)
	must(err)

	return GlitchyImage{
		path,
		pixels,
	}
}

// Corrupt is the process that changes the bytes in the image
func (gi *GlitchyImage) Corrupt(level int) {

}

// Save just saves the new corrupted image
func (gi *GlitchyImage) Save(path string) {

}

func must(e error) {
	if e != nil {
		panic(e)
	}
}
