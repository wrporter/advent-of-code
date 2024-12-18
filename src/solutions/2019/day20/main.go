package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/timeit"
	maze2 "aoc/src/solutions/2019/day20/lib/maze"
	"fmt"
	"time"
)

func main() {
	runMaze("./day20/sample1.txt", false)
	runMaze("./day20/sample2.txt", false)
	runMaze("./day20/sample3.txt", false)
	runMaze("./day20/input.txt", false)

	runMaze("./day20/sample1.txt", true)
	runMaze("./day20/sample2.txt", true)
	runMaze("./day20/sample3.txt", true)
	runMaze("./day20/input.txt", true)
}

func runMaze(inputFile string, recursive bool) {
	lines, _ := file.ReadFile(inputFile)
	defer timeit.Track(time.Now(), fmt.Sprintf("run maze (rec: %v): %s", recursive, inputFile))
	m := maze2.New(lines)
	fmt.Println(m.MinSteps(recursive))
}
