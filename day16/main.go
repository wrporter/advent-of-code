package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day16/internal/fft"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
)

func main() {
	lines, _ := file.ReadFile("./day16/input.txt")
	signal := lines[0]
	//signal := "03036732577212944063491565474664"

	fmt.Printf("Output signal: %s\n", fft.Apply(signal, 100)[:8])
	fmt.Printf("Decoded signal: %s\n", fft.Decode(signal))
}
