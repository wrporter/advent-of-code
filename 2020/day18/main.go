package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
)

func main() {
	year, day := 2020, 18
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	sum := 0
	for _, line := range input {
		overwritten := strings.ReplaceAll(line, "*", "-") // rewrite parser token precedence
		result := Evaluate(overwritten)
		sum += result
	}
	return sum
}

func part2(input []string) interface{} {
	sum := 0
	for _, line := range input {
		overwritten := strings.ReplaceAll(line, "*", "-")       // rewrite parser token precedence
		overwritten = strings.ReplaceAll(overwritten, "+", "/") // rewrite parser token precedence
		result := Evaluate(overwritten)
		sum += result
	}
	return sum
}

func Evaluate(expression string) int {
	exp, err := parser.ParseExpr(expression)
	if err != nil {
		fmt.Println(err)
	}
	return Eval(exp)
}

func Eval(exp ast.Expr) int {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.ParenExpr:
		return EvalParenExpr(exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, _ := strconv.Atoi(exp.Value)
			return i
		}
	}

	return 0
}

func EvalParenExpr(exp *ast.ParenExpr) int {
	return Eval(exp.X)
}

func EvalBinaryExpr(exp *ast.BinaryExpr) int {
	left := Eval(exp.X)
	right := Eval(exp.Y)

	switch exp.Op {
	case token.ADD:
		return left + right
	case token.QUO:
		return left + right
	case token.SUB:
		return left * right
	}

	return 0
}
