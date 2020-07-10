package main

import (
	"adrian/glitch"
	"fmt"
)

func main() {
	gi := glitch.NewGlitchyImage("image.jpeg")
	gi.Corrupt(10)
	gi.Save("corrupt.jpeg")
	fmt.Println(gi)
}
