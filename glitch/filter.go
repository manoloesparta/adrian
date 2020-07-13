package glitch

import (
	"math/rand"
	"time"
)

func choose(img []byte) {
	rand.Seed(time.Now().UTC().UnixNano())

	opt := rand.Intn(6)
	switch opt {
	case 0, 1, 2:
		modify(img)
	case 3:
		delete(img)
	case 4, 5, 6:
		swap(img)
	}
}

func modify(img []byte) {
	point := rand.Intn(len(img) - 8)
	for i := point; i < point+8; i++ {
		fact := rand.Intn(4) - 2
		img[i] = (img[i] + (byte)(fact)) % 255
	}
}

func delete(img []byte) {
	point := rand.Intn(len(img) - 8)
	img = append(img[:point], img[point+8:]...)
}

func swap(img []byte) {
	one := rand.Intn(len(img) - 8)
	sli1 := img[one : one+8]

	two := rand.Intn(len(img) - 8)
	sli2 := img[two : two+8]

	tmp1 := append(img[:one], sli2...)
	tmp2 := append(sli1, img[two:]...)

	img = append(tmp1, tmp2...)
}
