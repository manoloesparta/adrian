package glitch

import "math/rand"

// Filter represents the operations over image

func choose(img *GlitchyImage) {
	opt := rand.Intn(3)
	switch opt {
	case 0:
		modify(img.pixels)
	case 1:
		delete(img.pixels)
	case 2:
		swap(img.pixels)
	}
}

func modify(img []byte) {

}

func delete(img []byte) {

}

func swap(img []byte) {

}
