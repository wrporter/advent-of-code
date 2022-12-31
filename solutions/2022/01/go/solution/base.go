package solution

import "github.com/wrporter/advent-of-code/internal/common/v2/solution"

func Run() {
	s := New()
	s.Run([]interface{}{}, []interface{}{})
}

func New() Solution {
	s := Solution{AbstractSolution: solution.AbstractSolution{
		Solution: Solution{},
		Year:     2022,
		Day:      01,
	}}
	return s
}

type Solution struct {
	solution.AbstractSolution
}
