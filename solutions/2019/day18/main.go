package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	vault2 "github.com/wrporter/advent-of-code/solutions/2019/day18/internal/vault"
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
