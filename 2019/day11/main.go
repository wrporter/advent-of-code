package main

import (
	"fmt"
	paint2 "github.com/wrporter/advent-of-code/2019/day11/internal/paint"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"strings"
)

func main() {
	codeLines, _ := file.ReadFile("./2019/day11/input.txt")
	code, _ := convert.ToInts(strings.Split(codeLines[0], ","))
	robot := paint2.NewRobot(code)
	//numPaintedPanels, region := robot.Paint(paint.Black)
	//fmt.Println(numPaintedPanels)
	//fmt.Println(paint.RenderPaintedRegion(region))

	numPaintedPanels, region := robot.Paint(paint2.White)
	fmt.Println(numPaintedPanels)
	fmt.Println(paint2.RenderPaintedRegion(region))
}
