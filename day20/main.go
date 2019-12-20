package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day20/internal/maze"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"github.com/wrporter/advent-of-code-2019/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	runSampleMaze(`         A           
         A           
  #######.#########  
  #######.........#  
  #######.#######.#  
  #######.#######.#  
  #######.#######.#  
  #####  B    ###.#  
BC...##  C    ###.#  
  ##.##       ###.#  
  ##...DE  F  ###.#  
  #####    G  ###.#  
  #########.#####.#  
DE..#######...###.#  
  #.#########.###.#  
FG..#########.....#  
  ###########.#####  
             Z       
             Z       `)
	runSampleMaze(`
"                   A               "
"                   A               "
"  #################.#############  "
"  #.#...#...................#.#.#  "
"  #.#.#.###.###.###.#########.#.#  "
"  #.#.#.......#...#.....#.#.#...#  "
"  #.#########.###.#####.#.#.###.#  "
"  #.............#.#.....#.......#  "
"  ###.###########.###.#####.#.#.#  "
"  #.....#        A   C    #.#.#.#  "
"  #######        S   P    #####.#  "
"  #.#...#                 #......VT"
"  #.#.#.#                 #.#####  "
"  #...#.#               YN....#.#  "
"  #.###.#                 #####.#  "
"DI....#.#                 #.....#  "
"  #####.#                 #.###.#  "
"ZZ......#               QG....#..AS"
"  ###.###                 #######  "
"JO..#.#.#                 #.....#  "
"  #.#.#.#                 ###.#.#  "
"  #...#..DI             BU....#..LF"
"  #####.#                 #.#####  "
"YN......#               VT..#....QG"
"  #.###.#                 #.###.#  "
"  #.#...#                 #.....#  "
"  ###.###    J L     J    #.#.###  "
"  #.....#    O F     P    #.#...#  "
"  #.###.#####.#.#####.#####.###.#  "
"  #...#.#.#...#.....#.....#.#...#  "
"  #.#####.###.###.#.#.#########.#  "
"  #...#.#.....#...#.#.#.#.....#.#  "
"  #.###.#####.###.###.#.#.#######  "
"  #.#.........#...#.............#  "
"  #########.###.###.#############  "
"           B   J   C               "
"           U   P   P               "
`)
	runMaze("./day20/input.txt")
}

func runSampleMaze(input string) {
	lines := strings.Split(input, "\n")
	defer timeit.Track(time.Now(), "find ZZ")
	m := maze.New(lines)
	fmt.Println(m.MinSteps())
}

func runMaze(inputFile string) {
	lines, _ := file.ReadFile(inputFile)
	defer timeit.Track(time.Now(), "find ZZ")
	m := maze.New(lines)
	fmt.Println(m.MinSteps())
}
