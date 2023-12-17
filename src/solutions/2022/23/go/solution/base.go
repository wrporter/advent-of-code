package solution

import "aoc/src/lib/go/v2/solution"

func Run() {
	s := New()
	s.Run([]interface{}{}, []interface{}{})
}

func New() Solution {
	s := Solution{}
	s.AbstractSolution = solution.New(s, 2022, 23)
	return s
}

type Solution struct {
	*solution.AbstractSolution
}
