package main

import (
	"aoc/src/lib/go/file"
	fft2 "aoc/src/solutions/2019/day16/lib/fft"
	"fmt"
)

func main() {
	lines, _ := file.ReadFile("./2019/day16/input.txt")
	signal := lines[0]
	//signal := "03036732577212944063491565474664"

	fmt.Printf("Output signal: %s\n", fft2.Apply(signal, 100)[:8])
	fmt.Printf("Decoded signal: %s\n", fft2.Decode(signal))
}
