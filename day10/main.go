package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day10/internal/monitor"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
)

func main() {
	lines, _ := file.ReadFile("./day10/input.txt")
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
	field := monitor.SplitLines(lines)
	m := monitor.New()
	station, _ := m.FindBestAsteroid(field)
	fmt.Println(station)
	fmt.Println(m.ZapAsteroids(field, 200))
}
