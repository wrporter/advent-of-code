package main

import (
	"aoc/src/solutions/2020/day18/parser"
	"fmt"
)

func main() {
	result := parser.Evaluate("2 + 3 * (4 + 5)")
	fmt.Println(result)
}
