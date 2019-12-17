package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day16/internal/fft"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
)

func main() {
	lines, _ := file.ReadFile("./day16/input.txt")
	input := lines[0]
	//input := "12345678"

	const digits = 8
	const phases = 100

	output := fft.Apply(input, phases)[:digits]
	fmt.Printf("Output signal: %s\n", output)
}
