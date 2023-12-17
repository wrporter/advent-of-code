package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	paint2 "aoc/src/solutions/2019/day11/lib/paint"
	"fmt"
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
