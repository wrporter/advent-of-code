package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day18/internal/vault"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"log"
	"time"
)

func main() {
	runMaze("./day18/input.txt")
	runMaze("./day18/input2.txt")
}

func runMaze(inputFile string) {
	lines, _ := file.ReadFile(inputFile)
	defer timeTrack(time.Now(), "find keys")
	v := vault.New(lines)
	fmt.Println(v.MinSteps())
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
