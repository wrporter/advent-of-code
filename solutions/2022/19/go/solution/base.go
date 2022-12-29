package solution

import "github.com/wrporter/advent-of-code/internal/common/solution"

func Run() {
	s := New()
	s.Run([]interface{}{}, []interface{}{})
}

func New() Solution {
	s := Solution{AbstractSolution: solution.AbstractSolution{
		Solution: Solution{},
		Year:     2022,
		Day:      19,
	}}
	return s
}

type Solution struct {
	solution.AbstractSolution
}
