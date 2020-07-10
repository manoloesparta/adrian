package glitch

import "math/rand"

// Filter represents the operations over image
type Filter struct{}

func (f *Filter) choose(img *GlitchyImage) {
	opt := rand.Intn(3)
	switch opt {
	case 0:
		f.modify(img)
	case 1:
		f.delete(img)
	case 2:
		f.swap(img)
	}
}

func (f *Filter) modify(img *GlitchyImage) {

}

func (f *Filter) delete(img *GlitchyImage) {

}

func (f *Filter) swap(img *GlitchyImage) {

}
