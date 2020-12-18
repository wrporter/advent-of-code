package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/2020/day18/parser"
)

func main() {
	result := parser.Evaluate("2 + 3 * (4 + 5)")
	fmt.Println(result)
}
