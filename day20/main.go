package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day20/internal/maze"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"github.com/wrporter/advent-of-code-2019/internal/common/timeit"
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
	m := maze.New(lines)
	fmt.Println(m.MinSteps(recursive))
}
