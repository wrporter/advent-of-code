package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day11/internal/paint"
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"strings"
)

func main() {
	codeLines, _ := file.ReadFile("./day11/input.txt")
	code, _ := conversion.ToInts(strings.Split(codeLines[0], ","))
	robot := paint.NewRobot(code)
	numPaintedPanels := robot.Paint()
	fmt.Println(numPaintedPanels)
}
