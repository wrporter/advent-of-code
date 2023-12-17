package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/timeit"
	vault2 "aoc/src/solutions/2019/day18/lib/vault"
	"fmt"
	"time"
)

func main() {
	runMaze("./day18/input.txt")
	runMaze("./day18/input2.txt")
}

func runMaze(inputFile string) {
	lines, _ := file.ReadFile(inputFile)
	defer timeit.Track(time.Now(), "find keys")
	v := vault2.New(lines)
	fmt.Println(v.MinSteps())
}
