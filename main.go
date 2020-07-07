package main

import (
	"adrian/glitch"
	"fmt"
	_ "image/jpeg"
)

func main() {
	gi := glitch.NewGlitchyImage("image.jpeg")
	fmt.Println(gi)
}
