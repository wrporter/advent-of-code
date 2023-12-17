package solution

import "aoc/src/lib/go/v2/solution"

func Run() {
	s := New()
	s.Run([]interface{}{}, []interface{}{})
}

func New() Solution {
	s := Solution{AbstractSolution: solution.AbstractSolution{
		Solution: Solution{},
		Year:     2022,
		Day:      2,
	}}
	return s
}

type Solution struct {
	solution.AbstractSolution
}
