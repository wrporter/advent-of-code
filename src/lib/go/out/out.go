package out

import (
	"aoc/src/lib/go/out/color"
	"fmt"
)

func Day(year int, day int) {
	fmt.Printf("🎄 %s%s%d: Day %d\n%s", color.Green, color.Underlined, year, day, color.Reset)
}

func Part1(answer interface{}) {
	fmt.Printf("⭐  %sPart 1: %s%v\n%s", color.Green, color.Red, answer, color.Reset)
}

func Part2(answer interface{}) {
	fmt.Printf("⭐  %sPart 2: %s%v\n%s", color.Green, color.Red, answer, color.Reset)
}
