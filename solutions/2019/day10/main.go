package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	monitor2 "github.com/wrporter/advent-of-code/solutions/2019/day10/internal/monitor"
)

func main() {
	lines, _ := file.ReadFile("./2019/day10/input.txt")
	//lines := []string{
	//	".#..##.###...#######",
	//	"##.############..##.",
	//	".#.######.########.#",
	//	".###.#######.####.#.",
	//	"#####.##.#.##.###.##",
	//	"..#####..#.#########",
	//	"####################",
	//	"#.####....###.#.#.##",
	//	"##.#################",
	//	"#####.##.###..####..",
	//	"..######..##.#######",
	//	"####.##.####...##..#",
	//	".#####..#.######.###",
	//	"##...#.##########...",
	//	"#.##########.#######",
	//	".####.#.###.###.#.##",
	//	"....##.##.###..#####",
	//	".#.#.###########.###",
	//	"#.#.#.#####.####.###",
	//	"###.##.####.##.#..##",
	//}
	field := monitor2.SplitLines(lines)
	m := monitor2.New()
	station, _ := m.FindBestAsteroid(field)
	fmt.Println(station)
	fmt.Println(m.ZapAsteroids(field, 200))
}
