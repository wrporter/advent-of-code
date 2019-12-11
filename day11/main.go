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
	//numPaintedPanels, region := robot.Paint(paint.Black)
	//fmt.Println(numPaintedPanels)
	//fmt.Println(paint.RenderPaintedRegion(region))

	numPaintedPanels, region := robot.Paint(paint.White)
	fmt.Println(numPaintedPanels)
	fmt.Println(paint.RenderPaintedRegion(region))
}
