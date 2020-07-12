package main

import (
	"adrian/glitch"
)

func main() {
	gi := glitch.NewGlitchyImage("image.jpeg")
	gi.Corrupt(10)
	gi.Save("corrupt.jpeg")
}
