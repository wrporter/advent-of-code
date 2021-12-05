package sat

// SummedAreaTable is a data structure that takes O(n*m) space and O(1) time for summing rectangular windows of values
// in a grid of numbers. See https://en.wikipedia.org/wiki/Summed-area_table for more information.
type SummedAreaTable struct {
	Width  int
	Height int
	grid   [][]int
	sat    [][]int
}

func NewSummedAreaTable(grid [][]int) *SummedAreaTable {
	return &SummedAreaTable{
		Width:  len(grid[0]),
		Height: len(grid),
		grid:   grid,
		sat:    CreateSummedAreaTable(grid),
	}
}

func CreateSummedAreaTable(grid [][]int) [][]int {
	sat := make([][]int, len(grid))
	sat[0] = make([]int, len(grid[0]))

	// copy first row
	for x := range grid[0] {
		sat[0][x] = grid[0][x]
	}

	// sum columns
	for y := 1; y < len(grid); y++ {
		sat[y] = make([]int, len(grid[y]))
		for x := range grid[y] {
			sat[y][x] = grid[y][x] + sat[y-1][x]
		}
	}

	// sum rows
	for y := range grid {
		for x := 1; x < len(grid[y]); x++ {
			sat[y][x] += sat[y][x-1]
		}
	}

	return sat
}

// SumWindow returns the sum of grid values from the top left corner to the bottom left corner of the coordinates
// provided. The coordinate at (tlx, tly) represents the top-left corner while (brx, bry) provide the bottom-right
// corner.
func (s *SummedAreaTable) SumWindow(tlx, tly, brx, bry int) int {
	result := s.sat[bry][brx]

	if tly > 0 {
		result -= s.sat[tly-1][brx]
	}

	if tlx > 0 {
		result -= s.sat[bry][tlx-1]
	}

	if tly > 0 && tlx > 0 {
		result += s.sat[tly-1][tlx-1]
	}

	return result
}
