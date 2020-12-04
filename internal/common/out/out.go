package out

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/out/color"
)

func Day(year int, day int) {
	fmt.Printf("🎄 %s%s%d: Day %d\n%s", color.Green, color.Underlined, year, day, color.Reset)
}

func Part1(answer interface{}) {
	fmt.Printf("🎄 %sPart1: %s%v\n%s", color.Green, color.Red, answer, color.Reset)
}

func Part2(answer interface{}) {
	fmt.Printf("🎄 %sPart2: %s%v\n%s", color.Green, color.Red, answer, color.Reset)
}
