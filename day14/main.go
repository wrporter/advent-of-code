package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day14/internal/nanofactory"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
)

func main() {
	lines, _ := file.ReadFile("./day14/input.txt")
	//lines := []string{
	//	"2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG",
	//	"17 NVRVD, 3 JNWZP => 8 VPVL",
	//	"53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL",
	//	"22 VJHF, 37 MNCFX => 5 FWMGM",
	//	"139 ORE => 4 NVRVD",
	//	"144 ORE => 7 JNWZP",
	//	"5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC",
	//	"5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV",
	//	"145 ORE => 6 MNCFX",
	//	"1 NVRVD => 8 CXFTF",
	//	"1 VJHF, 6 MNCFX => 4 RFSQX",
	//	"176 ORE => 6 VJHF",
	//}
	//lines := []string{
	//	"10 ORE => 10 A",
	//	"1 ORE => 1 B",
	//	"7 A, 1 B => 1 C",
	//	"7 A, 1 C => 1 D",
	//	"7 A, 1 D => 1 E",
	//	"7 A, 1 E => 1 FUEL",
	//}
	factory := nanofactory.New()
	fmt.Println(factory.GetRequiredOre(lines, 1))
}
