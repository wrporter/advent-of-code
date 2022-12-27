package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	fft2 "github.com/wrporter/advent-of-code/solutions/2019/day16/internal/fft"
)

func main() {
	lines, _ := file.ReadFile("./2019/day16/input.txt")
	signal := lines[0]
	//signal := "03036732577212944063491565474664"

	fmt.Printf("Output signal: %s\n", fft2.Apply(signal, 100)[:8])
	fmt.Printf("Decoded signal: %s\n", fft2.Decode(signal))
}
