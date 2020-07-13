package main

import (
	"adrian/glitch"
	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("f", "", "Path of jpeg to corrupt. (Required)")
	output := flag.String("o", "output.jpeg", "The name of the corrupted file")
	levels := flag.Int("l", 3, "How cooked you want your corrupted jpeg")
	flag.Parse()

	if *file == "" {
		fmt.Println("\nadrian -f [input file] -l [levels] -o [output file] \n ")

		flag.PrintDefaults()
		os.Exit(1)
	}

	if _, err := os.Stat(*file); os.IsNotExist(err) {
		fmt.Println("File doesn't exist")
		os.Exit(1)
	}

	gi := glitch.NewGlitchyImage(*file)
	gi.Corrupt(*levels)
	gi.Save(*output)
}
